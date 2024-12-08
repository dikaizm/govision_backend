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
	var doctorResponse []*response.GetDoctorProfilePreview

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

	doctorResponse = []*response.GetDoctorProfilePreview{}

	for _, doctor := range doctors {
		doctorResponse = append(doctorResponse, &response.GetDoctorProfilePreview{
			UserID:         doctor.UserID,
			Name:           doctor.User.Name,
			Specialization: doctor.Specialization,
			Rating:         doctor.Rating,
			WorkYears:      doctor.WorkYears,
			City:           doctor.City,
			Province:       doctor.Province,
			Photo:          doctor.User.Photo,
		})
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch all doctors",
		Data:    doctorResponse,
	}, http.StatusOK)
}

func (c *DoctorController) View(w http.ResponseWriter, r *http.Request) {
	var doctorResponse *response.GetDoctorProfile

	profileID := helpers.UrlVars(r, "user_id")

	if profileID == "" {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  "User ID is required",
		}, http.StatusBadRequest)
		return
	}

	doctor, err := c.doctorService.GetProfile(profileID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	doctorResponse = &response.GetDoctorProfile{
		UserID:         doctor.UserID,
		Name:           doctor.User.Name,
		StrNo:          doctor.StrNo,
		Photo:          doctor.User.Photo,
		Specialization: doctor.Specialization,
		Institution:    doctor.Institution,
		City:           doctor.City,
		Province:       doctor.Province,
		Rating:         doctor.Rating,
		TotalPatient:   doctor.TotalPatient,
		WorkYears:      doctor.WorkYears,
		BioDesc:        doctor.BioDesc,
		Experiences:    []response.DoctorExperience{},
		Educations:     []response.DoctorEducation{},
		Schedules:      []response.DoctorSchedule{},
	}

	if doctor.Experiences != nil {
		for _, exp := range doctor.Experiences {
			doctorResponse.Experiences = append(doctorResponse.Experiences, response.DoctorExperience{
				Institution: exp.InstitutionName,
				StartDate:   exp.StartDate,
				EndDate:     exp.EndDate,
			})
		}
	}
	if doctor.Educations != nil {
		for _, edu := range doctor.Educations {
			doctorResponse.Educations = append(doctorResponse.Educations, response.DoctorEducation{
				University: edu.University,
				Major:      edu.Major,
				StartYear:  edu.StartYear,
				EndYear:    edu.EndYear,
			})
		}
	}
	if doctor.Schedules != nil {
		for _, sch := range doctor.Schedules {
			doctorResponse.Schedules = append(doctorResponse.Schedules, response.DoctorSchedule{
				Date: sch.Date,
			})
		}
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch doctor profile success",
		Data:    doctorResponse,
	}, http.StatusOK)
}

func (c *DoctorController) Profile(w http.ResponseWriter, r *http.Request) {
	userID := helpers.UrlVars(r, "user_id")

	doctor, err := c.doctorService.GetProfile(userID)
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

func (c *DoctorController) GetTimeSlots(w http.ResponseWriter, r *http.Request) {
	var timeSlotsResponse []response.GetDoctorTimeSlot

	userID := helpers.UrlVars(r, "user_id")
	date := helpers.UrlVars(r, "date")

	timeSlots, err := c.doctorService.GetTimeSlots(userID, date)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	timeSlotsResponse = []response.GetDoctorTimeSlot{}

	for _, ts := range timeSlots {
		timeSlotsResponse = append(timeSlotsResponse, response.GetDoctorTimeSlot{
			ID:        ts.ID,
			StartTime: ts.StartTime,
			EndTime:   ts.EndTime,
			IsBooked:  ts.IsBooked,
		})
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch time slots success",
		Data:    timeSlotsResponse,
	}, http.StatusOK)
}
