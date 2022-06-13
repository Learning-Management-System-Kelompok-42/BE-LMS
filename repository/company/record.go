package company

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/requestCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"gorm.io/gorm"
)

type Company struct {
	ID             string `gorm:"primaryKey,size:200"`
	Name           string
	Address        string
	Web            string `gorm:"size:250;uniqueIndex"`
	Sector         string
	Logo           string
	Users          []users.User                  `gorm:"foreignKey:CompanyID"`
	RequestCourses []requestCourse.RequestCourse `gorm:"foreignKey:CompanyID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
