package company

import (
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) domain.CompanyRepository {
	var companyRepository domain.CompanyRepository

<<<<<<< Updated upstream
	if dbCon.Driver == util.MySQL {
		companyRepository = NewCompanyRepository(dbCon.MySQL)
=======
	if dbCon.Driver == util.PostgreSQL {
		companyRepository = NewCompanyRepository(dbCon.PostgreSQL)
>>>>>>> Stashed changes
	}

	return companyRepository
}
