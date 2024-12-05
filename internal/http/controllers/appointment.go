package controllers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type AppointmentController struct {
	aptService service_intf.AppointmentService
}

func NewAppointmentController(aptService service_intf.AppointmentService) controller_intf.AppointmentController {
	return &AppointmentController{
		aptService: aptService,
	}
}

func (c *AppointmentController) Create(w http.ResponseWriter, r *http.Request) {
	req := request.CreateAppointment{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	// Validate request body
	if err := validate.Struct(&req); err != nil {
		helpers.FailedValidation(w, err)
		return
	}

	// Get current user
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	req.PatientUserID = currentUser.ID

	_, err = c.aptService.Create(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Create appointment success",
	}, http.StatusCreated)
}

func (c *AppointmentController) ViewAll(w http.ResponseWriter, r *http.Request) {
	var appointmentResponse []*response.GetAppointment

	// Query params
	var filter request.FilterViewAllAppointment
	filter.Range = r.URL.Query().Get("range")
	if filter.Range != "" {
		switch filter.Range {
		case "today", "week", "month":
		default:
			helpers.SendResponse(w, response.Response{
				Status: "error",
				Error:  "Invalid range query",
			}, http.StatusBadRequest)
			return
		}
	}

	// Get current user
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	apts, err := c.aptService.FindAllByPatient(currentUser.ID, &filter)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	appointmentResponse = make([]*response.GetAppointment, 0)

	for _, apt := range apts {
		appointmentResponse = append(appointmentResponse, &response.GetAppointment{
			ID:                   apt.ID,
			DoctorUserID:         apt.Doctor.UserID,
			DoctorName:           apt.Doctor.User.Name,
			DoctorSpecialization: apt.Doctor.Specialization,
			DoctorPhotoURL:       apt.Doctor.User.Photo,

			Date:      apt.TimeSlot.Schedule.Date,
			StartTime: apt.TimeSlot.StartTime,
			EndTime:   apt.TimeSlot.EndTime,

			Location: apt.Doctor.Institution,
		})
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Get appointments success",
		Data:    appointmentResponse,
	}, http.StatusOK)
}

func (c *AppointmentController) Confirm(w http.ResponseWriter, r *http.Request) {
	req := request.ConfirmAppointment{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	// Validate request body
	err := validate.Struct(&req)
	if err != nil {
		helpers.FailedValidation(w, err)
		return
	}

	aptID, err := helpers.StringToInt64(helpers.UrlVars(r, "apt_id"))
	if err != nil {
		helpers.FailedGetUrlVars(w, err, nil)
		return
	}

	if err := c.aptService.UpdateStatus(*aptID, req.Confirm); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Update appointment status success",
	}, http.StatusOK)
}
