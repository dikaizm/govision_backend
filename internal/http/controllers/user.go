package controllers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type UserController struct {
	userService service_intf.UserService
	authService service_intf.AuthService
}

func NewUserController(userService service_intf.UserService, authService service_intf.AuthService) controller_intf.UserController {
	return &UserController{
		userService: userService,
		authService: authService,
	}
}

func (c *UserController) View(w http.ResponseWriter, r *http.Request) {
	var userResponse response.GetUser

	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	user, err := c.userService.Get(currentUser.ID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	userResponse = response.GetUser{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role.RoleName,
		Photo:  user.Photo,
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch user",
		Data:    userResponse,
	}, http.StatusOK)
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

func (c *UserController) CreatePatientProfile(w http.ResponseWriter, r *http.Request) {
	req := &request.RegisterPatient{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	if err := c.authService.RegisterAsPatient(currentUser.ID, req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to create patient profile",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Patient profile created successfully",
	}, http.StatusCreated)
}
