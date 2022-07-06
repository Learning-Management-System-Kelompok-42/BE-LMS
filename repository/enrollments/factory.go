package enrollments

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) enrollments.EnrollmentRepository {
	var enrollmentRepo enrollments.EnrollmentRepository

	if dbCon.Driver == util.PostgreSQL {
		enrollmentRepo = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return enrollmentRepo
}
