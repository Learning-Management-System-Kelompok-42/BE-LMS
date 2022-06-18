package module

import (
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) module.ModuleRepository {
	var moduleRepositroy module.ModuleRepository

	if dbCon.Driver == util.PostgreSQL {
		moduleRepositroy = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return moduleRepositroy
}
