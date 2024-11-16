package controllers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
	"github.com/mitchellh/mapstructure"
)

type DoctorController struct {
	doctorService service_intf.DoctorService
}

func NewDoctorController(doctorService service_intf.DoctorService) controller_intf.DoctorController {
	return &DoctorController{
		doctorService: doctorService,
	}
}

func (c *DoctorController) ViewAll(w http.ResponseWriter, r *http.Request) {
	var filterQuery map[string]string
	if err := helpers.QueryDecoder(r, &filterQuery); err != nil {
		helpers.FailedParsingQuery(w, err)
		return
	}

	filter := &request.FilterAppointmentSchedule{}
	if err := mapstructure.Decode(filterQuery, filter); err != nil {
		helpers.FailedParsingQuery(w, err)
		return
	}

	doctors, err := c.doctorService.FindAll(filter)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch all doctors",
		Data:    doctors,
	}, http.StatusOK)
}

func (c *DoctorController) Profile(w http.ResponseWriter, r *http.Request) {
	doctorID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.FailedGetUrlVars(w, err, nil)
		return
	}

	doctor, err := c.doctorService.GetProfile(*doctorID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch doctor profile success",
		Data:    doctor,
	}, http.StatusOK)
}

func (c *DoctorController) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	req := []*request.CreateDoctorSchedule{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	for _, s := range req {
		err := validate.Struct(s)
		if err != nil {
			helpers.FailedValidation(w, err)
			return
		}
	}

	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	if err = c.doctorService.CreateSchedule(currentUser.ID, req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Schedule created",
	}, http.StatusCreated)
}
