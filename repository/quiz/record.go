package quiz

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/options"
	"gorm.io/gorm"
	"time"
)

type Quiz struct {
	ID        string `gorm:"primaryKey;size:200"`
	ModuleID  string `gorm:"size:200"`
	Title     string
	Question  string
	Options   []options.Option `gorm:"foreignKey:QuizID"`
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
