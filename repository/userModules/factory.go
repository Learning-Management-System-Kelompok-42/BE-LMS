package userModules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbConnection *util.DatabaseConnection) userModules.UserModulesRepository {
	var usermodules userModules.UserModulesRepository

	if dbConnection.Driver == util.PostgreSQL {
		usermodules = NewPostgreSQLRepository(dbConnection.PostgreSQL)
	}

	return usermodules
}
