package course

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/module"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specializationCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userCourse"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID                    string `gorm:"primaryKey;size:200"`
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
