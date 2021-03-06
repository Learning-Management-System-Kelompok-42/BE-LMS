package userModules

import (
	"time"
)

type Domain struct {
	ID        string
	UserID    string
	ModuleID  string
	CourseID  string
	Point     int32
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProggresCourse(id, userID, courseID, moduleID string, point int32, status bool) Domain {
	return Domain{
		ID:        id,
		UserID:    userID,
		ModuleID:  moduleID,
		CourseID:  courseID,
		Point:     point,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (old *Domain) UpdateProgress(point int32, status bool) Domain {
	return Domain{
		ID:        old.ID,
		UserID:    old.UserID,
		ModuleID:  old.ModuleID,
		CourseID:  old.CourseID,
		Point:     point,
		Status:    status,
		CreatedAt: old.CreatedAt,
		UpdatedAt: time.Now(),
	}
}
