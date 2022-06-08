package api

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller *users.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/v1/register", controller.UserV1Controller.Register)

	userV1 := e.Group("/v1/users")
	userV1.GET("/:id", controller.UserV1Controller.GetUserByID)
}
