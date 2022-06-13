package faq

import (
	"gorm.io/gorm"
	"time"
)

type Faq struct {
	ID        string `gorm:"primaryKey;size:200"`
	Question  string
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
