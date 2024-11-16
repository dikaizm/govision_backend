package controllers

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
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

	res := response.Response{
		Status:  "success",
		Message: "Detect fundus success",
		Data:    newFundus,
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *FundusController) ViewFundus(w http.ResponseWriter, r *http.Request) {
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

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "View fundus success",
		Data:    fundus,
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

func (c *FundusController) RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request) {

}

func (c *FundusController) VerifyFundusByDoctor(w http.ResponseWriter, r *http.Request) {
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
