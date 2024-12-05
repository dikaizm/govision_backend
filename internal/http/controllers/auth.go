package controllers

import (
	"fmt"
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type AuthController struct {
	authService service_intf.AuthService
}

func NewAuthController(authService service_intf.AuthService) controller_intf.AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req := request.Register{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	fmt.Println("req", req)

	// Validate the request body
	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Validation error",
			Error:   helpers.GetValidationErrors(err),
		}, http.StatusBadRequest)
		return
	}

	user, err := c.authService.Register(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to register user",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := response.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    user,
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := request.Login{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Validate the request body
	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Validation error",
			Error:   helpers.GetValidationErrors(err),
		}, http.StatusBadRequest)
		return
	}

	user, err := c.authService.Login(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to login user",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := response.Response{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    user,
	}

	helpers.SendResponse(w, res, http.StatusOK)
}
