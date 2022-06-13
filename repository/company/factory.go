package company

import (
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) domain.CompanyRepository {
	var companyRepository domain.CompanyRepository

	if dbCon.Driver == util.PostgreSQL {
		companyRepository = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return companyRepository
}
