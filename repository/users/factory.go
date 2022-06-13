package users

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbConnection *util.DatabaseConnection) users.UserRepository {
	var userRepository users.UserRepository

	if dbConnection.Driver == util.PostgreSQL {
		userRepository = NewUserRepository(dbConnection.PostgreSQL)
	}

	return userRepository
}
