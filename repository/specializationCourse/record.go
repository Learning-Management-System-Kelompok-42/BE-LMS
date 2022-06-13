package specializationCourse

import (
	"gorm.io/gorm"
	"time"
)

type SpecializationCourse struct {
	ID               string `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID         string `gorm:"size:200"`
	SpecializationID string `gorm:"size:200"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
