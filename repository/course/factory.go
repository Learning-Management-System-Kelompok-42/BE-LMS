package course

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) course.CourseRepository {
	var courseRepository course.CourseRepository

	if dbCon.Driver == util.PostgreSQL {
		courseRepository = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}

	return courseRepository
}
