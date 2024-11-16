package controllers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/dikaizm/govision_backend/internal/http/controllers"
	"github.com/golang/mock/gomock"
)

func TestViewAll(t *testing.T) {
	c := &controllers.DoctorController{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest("GET", "/appointment?start_hour=11:00", nil)
	w := httptest.NewRecorder()

	c.ViewAll(w, req)
	res := w.Result()

	defer res.Body.Close()
}
