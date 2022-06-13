package certificate

import (
	"gorm.io/gorm"
	"time"
)

type Certificate struct {
	ID        string `gorm:"primaryKey;size:200"`
	CourseID  string `gorm:"size:200"`
	UserID    string `gorm:"size:200"`
	Signature string
	Expired   time.Time
	CreatedAt time.Time //will be aliasing as Publish
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
