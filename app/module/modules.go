package modules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	companyController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company"
	userController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	companyService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	userService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	companyRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/company"
	userRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	//initiate dependency injection for user
	userPermitRepo := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewUserService(userPermitRepo)
	userPermitControllerV1 := userController.NewController(userPermitService)

	//initiate dependency injection for company
	companyPermitRepo := companyRepo.RepositoryFactory(dbCon)
	companyPermitService := companyService.NewCompanyService(companyPermitRepo)
	companyPermitControllerV1 := companyController.NewController(companyPermitService)

	controllers := api.Controller{
		UserV1Controller:    userPermitControllerV1,
		CompanyV1Controller: companyPermitControllerV1,
	}

	return controllers
}
