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

func (repo *postgreSQLRepository) FindByID(id string) (course course.Domain, err error) {
	var newCourse Course
	err = repo.db.Where("id = ?", id).First(&newCourse).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return course, exception.ErrNotFound
		}
		return course, exception.ErrNotFound
	}

	course = newCourse.ToDomain()

	return course, nil
}
