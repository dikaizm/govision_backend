package route_intf

import (
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
)

type Controllers struct {
	Auth        controller_intf.AuthController
	Article     controller_intf.ArticleController
	Appointment controller_intf.AppointmentController
	Fundus      controller_intf.FundusController
	User        controller_intf.UserController
	Doctor      controller_intf.DoctorController
}
