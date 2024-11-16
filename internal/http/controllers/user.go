package controllers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type UserController struct {
	userService service_intf.UserService
}

func NewUserController(userService service_intf.UserService) controller_intf.UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) ViewPatientProfile(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	profile, err := c.userService.GetProfilePatient(currentUser.ID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch patient profile",
		Data:    profile,
	}, http.StatusOK)
}

func (c *UserController) ViewDoctorProfile(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	profile, err := c.userService.GetProfileDoctor(currentUser.ID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch doctor profile",
		Data:    profile,
	}, http.StatusOK)
}
