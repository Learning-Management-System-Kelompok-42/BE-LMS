package requestCourse

import (
	"gorm.io/gorm"
	"time"
)

type RequestCourse struct {
	ID         string `gorm:"primaryKey,size:200"`
	CompanyID  string `gorm:"size:200"`
	NameCourse string
	Reason     string
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
