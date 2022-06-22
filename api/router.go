package api

import (
	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller           *users.Controller
	CompanyV1Controller        *company.Controller
	SpecializationV1Controller *specialization.Controller
	QuizV1Controller           *quiz.Controller
	ModuleV1Controller         *module.Controller
	CourseV1Controller         *course.Controller
	AuthV1Controller           *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	// Register User
	e.POST("/v1/user/register", controller.UserV1Controller.Register)
	// Register Company
	e.POST("/v1/company/register", controller.CompanyV1Controller.Register)
	// Login User and Company
	e.POST("/v1/auth/login", controller.AuthV1Controller.Login)
	// Get invitation link
	e.GET("/v1/specialization", controller.SpecializationV1Controller.GetInvitation)

	userV1 := e.Group("/v1/users")
	userV1.Use(m.JWTMiddleware(config))
	userV1.GET("/:id", controller.UserV1Controller.GetUserByID)

	courseV1 := e.Group("/v1/course")
	courseV1.Use(m.JWTMiddleware(config))
	courseV1.POST("", controller.CourseV1Controller.Register, m.CheckLevelAccess)

	quizV1 := e.Group("/v1/quiz")
	quizV1.GET("/:id", controller.QuizV1Controller.FindByID)
	quizV1.PUT("", controller.QuizV1Controller.Update)
	quizV1.POST("", controller.QuizV1Controller.Create)

	moduleV1 := e.Group("/v1/module")
	moduleV1.GET("/:id", controller.ModuleV1Controller.GetByID)
	moduleV1.PUT("", controller.ModuleV1Controller.Update)
	moduleV1.POST("", controller.ModuleV1Controller.Register)

	companyV1 := e.Group("/v1/admin")
	companyV1.Use(m.JWTMiddleware(config))
	companyV1.POST("/specialization", controller.SpecializationV1Controller.Register, m.CheckLevelAccess)
	companyV1.GET("/users", controller.UserV1Controller.GetAllUsers, m.CheckLevelAccess)
	companyV1.GET("/dashboard", controller.CompanyV1Controller.GetDashboard, m.CheckLevelAccess)

}
