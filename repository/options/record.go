package options

import (
	"gorm.io/gorm"
	"time"
)

type Option struct {
	ID        string `gorm:"primaryKey;size:200"`
	QuizID    string `gorm:"size:200"`
	Option    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
