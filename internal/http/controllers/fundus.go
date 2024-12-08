package controllers

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
	"gorm.io/gorm"
)

type FundusController struct {
	fundusService service_intf.FundusService
}

func NewFundusController(fundusService service_intf.FundusService) controller_intf.FundusController {
	return &FundusController{
		fundusService: fundusService,
	}
}

func (c *FundusController) DetectFundusImage(w http.ResponseWriter, r *http.Request) {
	// Request image file
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse form",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Get image file
	file, _, err := r.FormFile("fundus_image")
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get image file",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to read image file",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Get current user
	user, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get current user",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Encode image to base64
	req := request.DetectFundusImage{
		UserID:      user.ID,
		FundusImage: base64.StdEncoding.EncodeToString(fileBytes),
	}

	newFundus, err := c.fundusService.DetectImage(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Load the Asia/Jakarta timezone
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		helpers.FailedGetTimezone(w)
		return
	}

	// Ensure CreatedAt is in Asia/Jakarta timezone
	newFundus.CreatedAt = newFundus.CreatedAt.In(loc)
	newFundus.UpdatedAt = newFundus.UpdatedAt.In(loc)

	res := response.Response{
		Status:  "success",
		Message: "Detect fundus success",
		Data:    newFundus,
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *FundusController) ViewFundusHistory(w http.ResponseWriter, r *http.Request) {
	var fundusResponse []*response.ViewFundusHistory

	user, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get current user",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	fundus, err := c.fundusService.ViewFundusHistory(user.ID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		helpers.FailedGetTimezone(w)
		return
	}

	for _, f := range fundus {
		pathArray := strings.Split(f.ImgURL, "/")
		trimmedPath := pathArray[len(pathArray)-1]

		fundusResponse = append(fundusResponse, &response.ViewFundusHistory{
			ID:               f.ID,
			ImageUrl:         trimmedPath,
			VerifyStatus:     f.VerifyStatus,
			PredictedDisease: f.PredictedDisease,
			CreatedAt:        f.CreatedAt.In(loc),
			UpdatedAt:        f.UpdatedAt.In(loc),
			Feedbacks:        []response.FundusFeedback{},
		})
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "View fundus history success",
		Data:    fundusResponse,
	}, http.StatusOK)
}

func (c *FundusController) ViewFundus(w http.ResponseWriter, r *http.Request) {
	var fundusResponse *response.ViewFundusHistory

	id, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	fundus, err := c.fundusService.ViewFundus(*id)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	pathArray := strings.Split(fundus.ImgURL, "/")
	trimmedPath := pathArray[len(pathArray)-1]

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		helpers.FailedGetTimezone(w)
		return
	}

	fundusResponse = &response.ViewFundusHistory{
		ID:               fundus.ID,
		ImageUrl:         trimmedPath,
		VerifyStatus:     fundus.VerifyStatus,
		PredictedDisease: fundus.PredictedDisease,
		CreatedAt:        fundus.CreatedAt.In(loc),
		UpdatedAt:        fundus.UpdatedAt.In(loc),
		Feedbacks:        []response.FundusFeedback{},
	}

	if (len(fundus.Feedbacks)) > 0 {
		for _, f := range fundus.Feedbacks {
			fundusResponse.Feedbacks = append(fundusResponse.Feedbacks, response.FundusFeedback{
				ID:           f.ID,
				DoctorUserID: f.Doctor.User.ID,
				DoctorName:   f.Doctor.User.Name,
				Notes:        f.Notes,
				CreatedAt:    f.CreatedAt.In(loc),
				UpdatedAt:    f.UpdatedAt.In(loc),
			})
		}
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "View fundus success",
		Data:    fundusResponse,
	}, http.StatusOK)
}

func (c *FundusController) DeleteFundus(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err = c.fundusService.DeleteFundus(*id); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Delete fundus success",
	}, http.StatusOK)
}

func (c *FundusController) UpdateVerifyFundusByDoctor(w http.ResponseWriter, r *http.Request) {
	fundusID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	req := request.VerifyFundus{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = c.fundusService.VerifyFundusByDoctor(*fundusID, req.DoctorID, req.Status, req.Notes)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to verify fundus",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Verify fundus success",
	}, http.StatusOK)
}

func (c *FundusController) ViewFundusImage(w http.ResponseWriter, r *http.Request) {
	path := helpers.UrlVars(r, "path")

	fundusImage, err := c.fundusService.GetFundusImage(path)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to locate fundus image",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, fundusImage)
}

func (c *FundusController) ViewVerifiedFundus(w http.ResponseWriter, r *http.Request) {
	var fundusResponse *response.ViewVerifiedFundus

	user, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get current user",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	fundusResponse, err = c.fundusService.ViewVerifiedFundus(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendResponse(w, response.Response{
				Status:  "success",
				Message: "No verified fundus found",
			}, http.StatusOK)
			return
		}

		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	pathArray := strings.Split(fundusResponse.ImageUrl, "/")
	trimmedPath := pathArray[len(pathArray)-1]

	fundusResponse.ImageUrl = trimmedPath

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "View verified fundus success",
		Data:    fundusResponse,
	}, http.StatusOK)
}

func (c *FundusController) RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request) {
	fundusID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = c.fundusService.RequestVerifyFundusByPatient(*fundusID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to request verify fundus",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Request verify fundus success",
	}, http.StatusOK)
}
