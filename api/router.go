package api

import (
	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller           *users.Controller
	CompanyV1Controller        *company.Controller
	SpecializationV1Controller *specialization.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/v1/user/register", controller.UserV1Controller.Register)
	e.POST("/v1/company/register", controller.CompanyV1Controller.Register)
	e.POST("/v1/specialization/register", controller.SpecializationV1Controller.Register)

	userV1 := e.Group("/v1/users")
	userV1.Use(m.JWTMiddleware())
	userV1.GET("/:id", controller.UserV1Controller.GetUserByID)

	companyV1 := e.Group("/v1/admin")
	companyV1.Use(m.JWTMiddleware())
	companyV1.POST("/specialization", controller.SpecializationV1Controller.Register, m.CheckLevelAccess)

}
