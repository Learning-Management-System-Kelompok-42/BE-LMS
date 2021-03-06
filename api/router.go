package api

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/enrollments"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/requestFeat"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/upload"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	EnrollmentV1Controller     *enrollments.Controller
	UserV1Controller           *users.Controller
	CompanyV1Controller        *company.Controller
	SpecializationV1Controller *specialization.Controller
	QuizV1Controller           *quiz.Controller
	ModuleV1Controller         *module.Controller
	CourseV1Controller         *course.Controller
	UserModuleV1Controller     *userModules.Controller
	RequestFeatV1Controller    *requestFeat.Controller
	AuthV1Controller           *auth.Controller
	UploadV1Controller         *upload.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	// HTTPS redirect
	// e.Use(middleware.HTTPSRedirect())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Redirect docs to stoplight.io
	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://capstone-lms.stoplight.io/docs/docs-api-lms/qw1jjvklit1ja-api-lms-capstone-project")
	})

	// Register User
	e.POST("/v1/user/register", controller.UserV1Controller.Register)
	// Register Company
	e.POST("/v1/company/register", controller.CompanyV1Controller.Register)
	// Login User and Company
	e.POST("/v1/login", controller.AuthV1Controller.Login)
	// Get invitation link
	e.GET("/v1/invitation", controller.SpecializationV1Controller.GetInvitation)

	e.POST("/v1/image", controller.UploadV1Controller.UploadFile)

	// userV1 := e.Group("/v1/dashboard")
	// userV1.Use(m.JWTMiddleware(config))
	// userV1.GET("/:id", controller.UserV1Controller.GetUserByID)

	companyV1 := e.Group("/v1/company")
	companyV1.Use(m.JWTMiddleware(config))

	// Dashboard company
	companyV1.GET("/:companyID/dashboard", controller.CompanyV1Controller.GetDashboard, m.CheckLevelAccess) //need to update UI, priority 1

	// Dashboard specialization
	companyV1.GET("/:companyID/specialization", controller.SpecializationV1Controller.GetAllSpecialization, m.CheckLevelAccess)
	companyV1.GET("/:companyID/specialization/:specializationID", controller.SpecializationV1Controller.GetDetailSpecialization, m.CheckLevelAccess)
	companyV1.GET("/:companyID/specialization/generate", controller.SpecializationV1Controller.GenerateLinkInvitation, m.CheckLevelAccess)
	companyV1.POST("/:companyID/specialization/:specializationID/course", controller.SpecializationV1Controller.RegisterCourseSpecialization, m.CheckLevelAccess)
	companyV1.POST("/:companyID/specialization", controller.SpecializationV1Controller.Register, m.CheckLevelAccess)
	companyV1.PUT("/:companyID/specialization/:specializationID", controller.SpecializationV1Controller.UpdateSpecialization, m.CheckLevelAccess)

	// Dashboard courses
	companyV1.GET("/:companyID/course", controller.CourseV1Controller.GetAllCourseDashboard, m.CheckLevelAccess)
	companyV1.POST("/:companyID/course", controller.CourseV1Controller.RegisterCourse, m.CheckLevelAccess)
	companyV1.GET("/:companyID/course/:courseID", controller.CourseV1Controller.GetDetailCourseDashboard, m.CheckLevelAccess)
	companyV1.PUT("/:companyID/course/:courseID", controller.CourseV1Controller.UpdateCourse, m.CheckLevelAccess)
	companyV1.GET("/:companyID/course/:courseID/detail", controller.CourseV1Controller.GetCourseModules, m.CheckLevelAccess)

	// Dashboard employee
	companyV1.GET("/:companyID/employee", controller.UserV1Controller.GetAllUsers, m.CheckLevelAccess)
	companyV1.GET("/:companyID/employee/:employeeID", controller.UserV1Controller.GetDetailUserDashboard, m.CheckLevelAccess) //need to add progress course, priority 6
	companyV1.PUT("/:companyID/employee/:employeeID", controller.UserV1Controller.UpdateSpecializationName, m.CheckLevelAccess)
	companyV1.PUT("/:companyID/employee/:employeeID/profile", controller.UserV1Controller.UpdateProfile)
	companyV1.PUT("/:companyID/employee/:employeeID/password", controller.UserV1Controller.ChangePassword)

	// Dashboard setting
	companyV1.PUT("/:companyID", controller.CompanyV1Controller.UpdateCompanyProfile, m.CheckLevelAccess) //add priority 9
	companyV1.GET("/:companyID", controller.CompanyV1Controller.GetCompanyProfile, m.CheckLevelAccess)
	// companyV1.GET("/:companyID/setting/:employeeID", controller.CompanyV1Controller.Profile, m.CheckLevelAccess) //add priority 8

	employeeV1 := e.Group("/v1/employee")
	employeeV1.Use(m.JWTMiddleware(config))

	// Dashboard employee
	employeeV1.GET("/:employeeID/dashboard/:specializationID", controller.UserV1Controller.GetDashboard)
	employeeV1.GET("/:employeeID", controller.UserV1Controller.GetDetailUsersByID)

	// Dashboard course
	employeeV1.GET("/:employeeID/course/:specializationID", controller.CourseV1Controller.GetAllCourse)
	employeeV1.GET("/:employeeID/course/:courseID/details", controller.CourseV1Controller.GetDetailCourse)

	// endpoint for modules course user
	employeeV1.GET("/:employeeID/course/:courseID/modules", controller.ModuleV1Controller.GetAllModules)

	// Enroll course
	employeeV1.POST("/:employeeID/course/:courseID/enroll", controller.EnrollmentV1Controller.CreateEnrollments)
	employeeV1.PUT("/:employeeID/course/:courseID/feedback", controller.EnrollmentV1Controller.CreateRatingReviews)

	// Proggress course
	employeeV1.PUT("/:employeeID/course/:courseID/progress", controller.UserModuleV1Controller.UpdateProgress)

	// Request Conseling And Course
	employeeV1.POST("/:employeeID/request", controller.RequestFeatV1Controller.CreateRequestFeat)
}
