package modules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	authController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth"
	companyController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	specializationController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization"
	userController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	authService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"
	companyService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	specializationService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	userService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	authRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/auth"
	companyRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/company"
	specializationRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"
	userRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	//initiate dependency injection for user
	userPermitRepo := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewUserService(userPermitRepo)
	userPermitControllerV1 := userController.NewController(userPermitService)

	//initiate dependency injection for company
	companyPermitRepo := companyRepo.RepositoryFactory(dbCon)
	companyPermitService := companyService.NewCompanyService(companyPermitRepo, userPermitRepo)
	companyPermitControllerV1 := companyController.NewController(companyPermitService)

	//initiate dependency injection for specialization
	specializationPermitRepo := specializationRepo.RepositoryFactory(dbCon)
	specializationPermitService := specializationService.NewSpecializationService(specializationPermitRepo)
	specializationPermitControllerV1 := specializationController.NewController(specializationPermitService)

	// initiate dependency injection for auth
	authPermitRepo := authRepo.RepositoryFactory(dbCon)
	authPermitService := authService.NewAuthService(authPermitRepo, config)
	authPermitControllerV1 := authController.NewController(authPermitService)

	controllers := api.Controller{
		UserV1Controller:           userPermitControllerV1,
		CompanyV1Controller:        companyPermitControllerV1,
		SpecializationV1Controller: specializationPermitControllerV1,
		AuthV1Controller:           authPermitControllerV1,
	}

	return controllers
}
