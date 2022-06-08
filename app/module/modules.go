package modules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api"
	userController "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users"
	userService "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	userRepo "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	//initiate Register user
	userPermitRepo := userRepo.RepositoryFactory(dbCon)
	userPermitService := userService.NewUserService(userPermitRepo)
	userPermitControllerV1 := userController.NewController(userPermitService)

	controllers := api.Controller{
		UserV1Controller: userPermitControllerV1,
	}

	return controllers
}
