package helpers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/response"
	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	errorsMap := make(map[string]string)
	for _, err := range validationErrors {
		errorsMap[err.Field()] = err.Tag()
	}

	return errorsMap
}

func FailedParsingBody(w http.ResponseWriter, err error) {
	SendResponse(w, response.Response{
		Status:  "error",
		Message: "Failed to parse request body",
		Error:   err.Error(),
	}, http.StatusBadRequest)
}

func FailedParsingQuery(w http.ResponseWriter, err error) {
	SendResponse(w, response.Response{
		Status:  "error",
		Message: "Failed to parse query",
		Error:   err.Error(),
	}, http.StatusBadRequest)
}

func FailedGetCurrentUser(w http.ResponseWriter, err error) {
	SendResponse(w, response.Response{
		Status:  "error",
		Message: "Failed to get current user",
		Error:   err.Error(),
	}, http.StatusBadRequest)
}

func FailedGetUrlVars(w http.ResponseWriter, err error, msg *string) {
	defaultMsg := "Failed to get url vars"
	if msg == nil {
		msg = &defaultMsg
	}

	SendResponse(w, response.Response{
		Status:  "error",
		Message: *msg,
		Error:   err.Error(),
	}, http.StatusBadRequest)
}

func FailedValidation(w http.ResponseWriter, err error) {
	SendResponse(w, response.Response{
		Status:  "error",
		Message: "Validation error",
		Error:   GetValidationErrors(err),
	}, http.StatusBadRequest)
}

func FailedGetTimezone(w http.ResponseWriter) {
	SendResponse(w, response.Response{
		Status:  "error",
		Message: "Failed to get timezone",
	}, http.StatusInternalServerError)
}
