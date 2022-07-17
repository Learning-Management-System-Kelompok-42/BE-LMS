package userModules

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"
	"gorm.io/gorm"
)

type UserModule struct {
	ID        string `gorm:"primaryKey;size:200;autoIncrement"`
	UserID    string `gorm:"size:200"`
	ModuleID  string `gorm:"size:200"`
	CourseID  string
	Point     int32
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain userModules.Domain) UserModule {
	return UserModule{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ModuleID:  domain.ModuleID,
		CourseID:  domain.CourseID,
		Point:     domain.Point,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (userModule *UserModule) ToDomain() *userModules.Domain {
	return &userModules.Domain{
		ID:        userModule.ID,
		UserID:    userModule.UserID,
		ModuleID:  userModule.ModuleID,
		CourseID:  userModule.CourseID,
		Point:     userModule.Point,
		Status:    userModule.Status,
		CreatedAt: userModule.CreatedAt,
		UpdatedAt: userModule.UpdatedAt,
	}
}
