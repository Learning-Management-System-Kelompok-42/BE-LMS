package users

import (
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbConnection *util.DatabaseConnection) domain.UserRepository {
	var userRepository domain.UserRepository

	if dbConnection.Driver == util.MySQL {
		userRepository = NewUserRepository(dbConnection.MySQL)
	}

	return userRepository
}
