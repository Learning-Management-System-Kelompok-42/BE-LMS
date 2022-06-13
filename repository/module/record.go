package module

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/material"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/userModules"
	"gorm.io/gorm"
	"time"
)

type Module struct {
	ID          string `gorm:"primaryKey;size:200"`
	CourseID    string `gorm:"size:200"`
	Title       string
	Orders      int32
	UserModules []userModules.UserModule `gorm:"foreignKey:ModuleID"`
	Materials   []material.Material      `gorm:"foreignKey:ModuleID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
