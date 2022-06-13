package material

import (
	"gorm.io/gorm"
	"time"
)

type Material struct {
	ID        string `gorm:"primaryKey;size:200"`
	ModuleID  string `gorm:"size:200"`
	Title     string
	Url       string
	Type      string
	Orders    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
