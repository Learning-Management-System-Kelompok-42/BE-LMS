package api

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller    *users.Controller
	CompanyV1Controller *company.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/v1/user/register", controller.UserV1Controller.Register)
	e.POST("/v1/company/register", controller.CompanyV1Controller.Register)

	userV1 := e.Group("/v1/users")
	userV1.GET("/:id", controller.UserV1Controller.GetUserByID)
}
