package specialization

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) specialization.SpecializationRepository {
	var specializationRepository specialization.SpecializationRepository

	if dbCon.Driver == util.PostgreSQL {
		specializationRepository = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return specializationRepository
}
