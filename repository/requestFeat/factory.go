package requestFeat

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbConnection *util.DatabaseConnection) requestFeat.RequestFeatRepository {
	var requestFeats requestFeat.RequestFeatRepository

	if dbConnection.Driver == util.PostgreSQL {
		requestFeats = NewPostgreSQLRepository(dbConnection.PostgreSQL)
	}

	return requestFeats
}
