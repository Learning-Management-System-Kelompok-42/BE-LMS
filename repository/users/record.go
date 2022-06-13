package users

import (
	// . "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/certificate"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userModules"
	"gorm.io/gorm"
)

type User struct {
	ID               string `gorm:"primaryKey;size:200;not null"`
	CompanyID        string `gorm:"size:200"`
	SpecializationID string `gorm:"size:200"`
	FullName         string
	Email            string `gorm:"size:250;uniqueIndex"`
	Password         string
	PhoneNumber      string
	Address          string
	Role             string
	LevelAccess      string
	UserCourses      []userCourse.UserCourse   `gorm:"foreignKey:CourseID"`
	UserModules      []userModules.UserModule  `gorm:"foreignKey:CourseID"`
	Certificates     []certificate.Certificate `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		ID:               user.ID,
		CompanyID:        user.CompanyID,
		Role:             user.Role,
		SpecializationID: user.SpecializationID,
		FullName:         user.FullName,
		Email:            user.Email,
		Password:         user.Password,
		PhoneNumber:      user.PhoneNumber,
		Address:          user.Address,
		LevelAccess:      user.LevelAccess,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:               domain.ID,
		CompanyID:        domain.CompanyID,
		SpecializationID: domain.SpecializationID,
		FullName:         domain.FullName,
		Email:            domain.Email,
		Password:         domain.Password,
		PhoneNumber:      domain.PhoneNumber,
		Address:          domain.Address,
		Role:             domain.Role,
		LevelAccess:      domain.LevelAccess,
		UserCourses:      nil,
		UserModules:      nil,
		Certificates:     nil,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}
