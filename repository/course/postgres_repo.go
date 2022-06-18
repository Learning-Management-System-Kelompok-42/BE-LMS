package course

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) course.CourseRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(course course.Domain) (id string, err error) {
	newCourse := FromDomain(course)
	err = repo.db.Create(&newCourse).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newCourse.ID

	return id, nil
}
