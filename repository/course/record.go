package course

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specializationCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userCourse"

	"gorm.io/gorm"
)

type Course struct {
	ID                    string `gorm:"primaryKey;size:200"`
	CompanyID             string `gorm:"size:200"`
	Title                 string
	Thumbnail             string
	Description           string
	SpecializationCourses []specializationCourse.SpecializationCourse `gorm:"primaryKey:CourseID"`
	Certificates          []certificate.Certificate                   `gorm:"primaryKey:CourseID"`
	Modules               []module.Module                             `gorm:"primaryKey:CourseID"`
	UserCourses           []userCourse.UserCourse                     `gorm:"primaryKey:CourseID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

func (courses *Course) ToDomain() course.Domain {
	return course.Domain{
		ID:          courses.ID,
		CompanyID:   courses.CompanyID,
		Title:       courses.Title,
		Thumbnail:   courses.Thumbnail,
		Description: courses.Description,
		CreatedAt:   courses.CreatedAt,
		UpdatedAt:   courses.UpdatedAt,
	}
}

func FromDomain(course course.Domain) Course {
	return Course{
		ID:          course.ID,
		CompanyID:   course.CompanyID,
		Title:       course.Title,
		Thumbnail:   course.Thumbnail,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
		DeletedAt:   gorm.DeletedAt{},
	}
}

func ToBatchList(courses []Course) []course.Domain {
	var coursesDomain []course.Domain
	for _, course := range courses {
		coursesDomain = append(coursesDomain, course.ToDomain())
	}
	return coursesDomain
}
