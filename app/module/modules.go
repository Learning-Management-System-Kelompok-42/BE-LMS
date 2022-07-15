package modules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	authController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth"
	companyController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	courseController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course"
	enrollmentController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/enrollments"
	moduleController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules"
	quizController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/quiz"
	specializationController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/upload"
	userController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	authService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"
	companyService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	courseService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	enrollmentService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	moduleService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	quizService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	specializationService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	userService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	authRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/auth"
	companyRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/company"
	courseRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/course"
	enrollmentRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/enrollments"
	moduleRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/modules"
	quizRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/quiz"
	specializationRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"
	userRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	// Initiate dependency injection for userModule
	userModulePermitRepo := userModuleRepo.RepositoryFactory(dbCon)
	userModulePermitService := userModuleService.NewUserModulesService(userModulePermitRepo)
	userModulePermitControllerV1 := userModuleController.NewController(userModulePermitService)

	// initiate dependency injection for enrollment
	enrollmentPermitRepo := enrollmentRepo.RepositoryFactory(dbCon)
	enrollmentPermitService := enrollmentService.NewEnrollmentService(enrollmentPermitRepo, userModulePermitRepo)
	enrollmentPermitControllerV1 := enrollmentController.NewController(enrollmentPermitService)

	// initiate dependency injection for quiz
	quizPermitRepo := quizRepo.RepositoryFactory(dbCon)
	quizPermitService := quizService.NewQuizService(quizPermitRepo)
	quizPermitControllerV1 := quizController.NewController(quizPermitService)

	// initiate dependency injection for modules
	modulePermitRepo := moduleRepo.RepositoryFactory(dbCon)
	modulePermitService := moduleService.NewModuleService(modulePermitRepo)
	modulePermitControllerV1 := moduleController.NewController(modulePermitService)

	//initiate dependency injection for user
	userPermitRepo := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewUserService(userPermitRepo)
	userPermitControllerV1 := userController.NewController(userPermitService)

	// initiate dependency injection for course
	coursePermitRepo := courseRepo.RepositoryFactory(dbCon)
	coursePermitService := courseService.NewCourseService(
		coursePermitRepo,
		userPermitRepo,
		enrollmentPermitRepo,
		modulePermitService,
		quizPermitService,
	)
	coursePermitControllerV1 := courseController.NewController(coursePermitService)

	//initiate dependency injection for company
	companyPermitRepo := companyRepo.RepositoryFactory(dbCon)
	companyPermitService := companyService.NewCompanyService(companyPermitRepo, userPermitRepo)
	companyPermitControllerV1 := companyController.NewController(companyPermitService)

	//initiate dependency injection for specialization
	specializationPermitRepo := specializationRepo.RepositoryFactory(dbCon)
	specializationPermitService := specializationService.NewSpecializationService(specializationPermitRepo, coursePermitRepo, userPermitRepo)
	specializationPermitControllerV1 := specializationController.NewController(specializationPermitService)

	// initiate dependency injection for auth
	authPermitRepo := authRepo.RepositoryFactory(dbCon)
	authPermitService := authService.NewAuthService(authPermitRepo, config)
	authPermitControllerV1 := authController.NewController(authPermitService)

	// Initiate dependency injection for upload
	uploadPermitControllerV1 := upload.NewController()

<<<<<<< Updated upstream
=======
	// Initiate dependency injection for RequestFeat
	requestFeatPermitRepo := requestFeatRepo.RepositoryFactory(dbCon)
	requestFeatPermitService := requestFeatService.NewRequestFeatService(requestFeatPermitRepo)
	requestFeatPermitControllerV1 := requestFeatController.NewController(requestFeatPermitService)

>>>>>>> Stashed changes
	controllers := api.Controller{
		UserV1Controller:           userPermitControllerV1,
		EnrollmentV1Controller:     enrollmentPermitControllerV1,
		CompanyV1Controller:        companyPermitControllerV1,
		SpecializationV1Controller: specializationPermitControllerV1,
		QuizV1Controller:           quizPermitControllerV1,
		ModuleV1Controller:         modulePermitControllerV1,
		CourseV1Controller:         coursePermitControllerV1,
		AuthV1Controller:           authPermitControllerV1,
		UploadV1Controller:         uploadPermitControllerV1,
	}

	return controllers
}
