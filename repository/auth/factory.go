package auth

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) auth.AuthRepository {
	var authRepository auth.AuthRepository

	if dbCon.Driver == util.PostgreSQL {
		authRepository = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return authRepository
}
