package userCourse

import (
	"time"

	"gorm.io/gorm"
)

type UserCourse struct {
	ID        string  `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID  string  `gorm:"size:200"`
	UserID    string  `gorm:"size:200"`
	Rating    float32 `gorm:"type:numeric(2,2)"`
	Reviews   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
