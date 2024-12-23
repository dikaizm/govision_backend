package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	"github.com/dikaizm/govision_backend/pkg/domain"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
	"gorm.io/gorm"
)

type ApiRequestBody struct {
	FundusImage string `json:"fundus_image"`
}

type ApiResponseData struct {
	PredictedClass string `json:"predicted_class"`
	CroppedImage   string `json:"cropped_image"`
}

type ApiResponseBody struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Error   string          `json:"error,omitempty"`
	Data    ApiResponseData `json:"data,omitempty"`
}

type FundusService struct {
	mlApi      string
	mlApiKey   string
	fundusRepo repo_intf.FundusRepository
	userRepo   repo_intf.UserRepository
}

func NewFundusService(mlApi string, mlApiKey string, fundusRepo repo_intf.FundusRepository, userRepo repo_intf.UserRepository) service_intf.FundusService {
	return &FundusService{
		mlApi:      mlApi,
		mlApiKey:   mlApiKey,
		fundusRepo: fundusRepo,
		userRepo:   userRepo,
	}
}

func detectFundusImageAPI(mlApi string, mlApiKey string, imageBlob string) (*ApiResponseBody, error) {
	// Create the request body
	requestBody, err := json.Marshal(ApiRequestBody{FundusImage: imageBlob})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/predict", mlApi), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", mlApiKey)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is not 200
	if resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response body
	var responseBody ApiResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	log.Printf("Predicted class: %s", responseBody.Data.PredictedClass)

	return &responseBody, nil
}

func (u *FundusService) DetectImage(p *request.DetectFundusImage) (res *response.DetectFundusImage, err error) {
	var imagePath string = ""

	patient, err := u.userRepo.FindPatientProfileByID(p.UserID)
	if err != nil {
		return nil, err
	}

	// Call machine learning API to detect fundus image
	mlResponse, err := detectFundusImageAPI(u.mlApi, u.mlApiKey, p.FundusImage)
	if err != nil {
		return nil, err
	}

	if (mlResponse.Data.PredictedClass != "") || (mlResponse.Data.CroppedImage != "") {
		// Store image in VM
		imagePath, err = helpers.StoreImage(mlResponse.Data.CroppedImage, "fundus")
		if err != nil {
			return nil, errors.New("failed to store image")
		}
	}

	// Map predicted class to domain constant
	switch mlResponse.Data.PredictedClass {
	case "No DR":
		mlResponse.Data.PredictedClass = domain.FundusDiseaseNoDR
	case "Mild":
		mlResponse.Data.PredictedClass = domain.FundusDiseaseMild
	case "Moderate":
		mlResponse.Data.PredictedClass = domain.FundusDiseaseModerate
	case "Severe":
		mlResponse.Data.PredictedClass = domain.FundusDiseaseSevere
	case "Proliferate DR":
		mlResponse.Data.PredictedClass = domain.FundusDiseaseProliferate
	default:
		mlResponse.Data.PredictedClass = domain.FundusDiseaseNotDetected
	}

	// Create fundus record in database
	fundus := &domain.CreateFundus{
		PatientID:        patient.ID,
		ImgURL:           imagePath,
		VerifyStatus:     domain.FundusVerifyStatusPending,
		PredictedDisease: mlResponse.Data.PredictedClass,
	}

	newFundus, err := u.fundusRepo.Create(fundus)
	if err != nil {
		return nil, err
	}

	res = &response.DetectFundusImage{
		ID:               newFundus.ID,
		VerifyStatus:     newFundus.VerifyStatus,
		PredictedDisease: newFundus.PredictedDisease,
		ImageBase64:      mlResponse.Data.CroppedImage,
		CreatedAt:        newFundus.CreatedAt,
	}

	return res, nil
}

func (u *FundusService) ViewFundus(fundusID int64) (*domain.Fundus, error) {
	fundus, err := u.fundusRepo.FindByID(fundusID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, errors.New("failed to find fundus record")
	}

	return fundus, nil
}

func (u *FundusService) ViewFundusHistory(userID string) ([]*domain.Fundus, error) {
	var fundusList []*domain.Fundus

	patient, err := u.userRepo.FindPatientProfileByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to find patient")
	}

	if patient != nil {
		fundusList, err = u.fundusRepo.FindAllByPatient(patient.ID)
		if err != nil {
			return nil, errors.New("failed to find fundus records")
		}
	}

	return fundusList, nil
}

func (u *FundusService) RequestVerifyFundusByPatient(fundusID int64) error {
	fundus, err := u.fundusRepo.FindByID(fundusID)
	if err != nil {
		return errors.New("failed to find fundus")
	}

	if fundus.VerifyStatus != domain.FundusVerifyStatusPending {
		return errors.New("fundus is not in pending status")
	}

	if err := u.fundusRepo.RequestVerifyStatusByPatient(fundus.ID); err != nil {
		return errors.New("failed to request verify fundus")
	}

	return nil
}

func (u *FundusService) VerifyFundusByDoctor(fundusID int64, doctorID int64, status string, notes string) error {
	if err := u.fundusRepo.UpdateVerifyStatusByDoctor(fundusID, doctorID, status); err != nil {
		return errors.New("failed to verify fundus")
	}

	if err := u.fundusRepo.CreateFeedbackByDoctor(fundusID, doctorID, notes); err != nil {
		return errors.New("failed storing feedbacks")
	}

	return nil
}

func (u *FundusService) DeleteFundus(fundusID int64) error {
	if err := u.fundusRepo.DeleteByID(fundusID); err != nil {
		return err
	}
	return nil
}

func (u *FundusService) GetFundusImage(path string) (string, error) {
	image, err := helpers.GetImageByPath("fundus", path)
	if err != nil {
		return "", errors.New("failed to get image")
	}
	return *image, nil
}

func (u *FundusService) ViewVerifiedFundus(userID string) (*response.ViewVerifiedFundus, error) {
	var fundusResponse *response.ViewVerifiedFundus

	// Get patient profile
	patient, err := u.userRepo.FindPatientProfileByID(userID)
	if err != nil {
		return nil, errors.New("failed to find patient")
	}

	// Get last verified fundus record
	verifiedFundus, err := u.fundusRepo.FindLastVerifiedByPatient(patient.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, errors.New("failed to find verified fundus")
	}

	// verifiedFundus
	fundusResponse = &response.ViewVerifiedFundus{
		ID:                     verifiedFundus.ID,
		ImageUrl:               verifiedFundus.ImgURL,
		VerifyStatus:           verifiedFundus.VerifyStatus,
		PredictedDisease:       verifiedFundus.PredictedDisease,
		DiabetesType:           patient.DiabetesType,
		RecommendedExamination: patient.RecommendedExamination,
		RecommendedNotes:       patient.RecommendedNotes,
	}

	return fundusResponse, nil
}
