package api

import (
	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller           *users.Controller
	CompanyV1Controller        *company.Controller
	SpecializationV1Controller *specialization.Controller
	AuthV1Controller           *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	e.POST("/v1/user/register", controller.UserV1Controller.Register)
	e.POST("/v1/company/register", controller.CompanyV1Controller.Register)
	e.POST("/v1/specialization/register", controller.SpecializationV1Controller.Register)
	e.POST("/v1/user/login", controller.AuthV1Controller.Login)
	e.GET("/v1/specialization", controller.SpecializationV1Controller.GetInvitation)

	userV1 := e.Group("/v1/users")
	userV1.Use(m.JWTMiddleware(config))
	userV1.GET("/:id", controller.UserV1Controller.GetUserByID)

	companyV1 := e.Group("/v1/admin")
	companyV1.Use(m.JWTMiddleware(config))
	companyV1.POST("/specialization", controller.SpecializationV1Controller.Register, m.CheckLevelAccess)

}
