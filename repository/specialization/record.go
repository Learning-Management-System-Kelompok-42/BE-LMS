package specialization

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specializationCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"gorm.io/gorm"
)

type Specialization struct {
	ID                    string `gorm:"primaryKey;size:200"`
	Name                  string
	Invitation            string
	Users                 []users.User                                `gorm:"foreignKey:SpecializationID"`
	SpecializationCourses []specializationCourse.SpecializationCourse `gorm:"foreignKey:SpecializationID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
