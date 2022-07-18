package course

import (
	"fmt"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/enrollments"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"

	"gorm.io/gorm"
)

type Course struct {
	ID                    string `gorm:"primaryKey;size:200"`
	CompanyID             string `gorm:"size:200"`
	Title                 string
	Thumbnail             string
	Description           string
	SpecializationCourses []specialization.SpecializationCourse `gorm:"primaryKey:CourseID"`
	Certificates          []certificate.Certificate             `gorm:"primaryKey:CourseID"`
	Modules               []module.Module                       `gorm:"primaryKey:CourseID"`
	Enrollments           []enrollments.Enrollments             `gorm:"primaryKey:CourseID"`
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

func (courses *Course) ToDomainInBatch() course.DomainCourseResp {
	return course.DomainCourseResp{
		ID:          courses.ID,
		CompanyID:   courses.CompanyID,
		Title:       courses.Title,
		Thumbnail:   courses.Thumbnail,
		Description: courses.Description,
		Modules:     module.ToDomainInBatch(courses.Modules),
	}
}

func ToBatchCourses(courses []Course) []course.DomainCourseResp {
	var coursesDomain []course.DomainCourseResp
	for _, course := range courses {
		fmt.Println("courses bathc = ", course)
		coursesDomain = append(coursesDomain, course.ToDomainInBatch())
	}
	return coursesDomain
}

func ToBatchList(courses []Course) []course.Domain {
	var coursesDomain []course.Domain
	for _, course := range courses {
		coursesDomain = append(coursesDomain, course.ToDomain())
	}
	return coursesDomain
}
