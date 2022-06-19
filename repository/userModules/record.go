package userModules

import (
	"time"

	"gorm.io/gorm"
)

type UserModule struct {
	ID        string `gorm:"primaryKey;size:200;autoIncrement"`
	UserID    string `gorm:"size:200"`
	ModuleID  string `gorm:"size:200"`
	Point     int32
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
