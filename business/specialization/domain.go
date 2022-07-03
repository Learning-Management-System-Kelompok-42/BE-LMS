package specialization

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
)

type Domain struct {
	ID         string
	CompanyID  string
	Name       string
	Invitation string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SpecializationDashboard struct {
	SpecializationID   string
	SpecializationName string
	AmountEmployee     int64
	AmountCourse       int64
}

type SpecializationDetail struct {
	SpecializationID   string
	CompanyID          string
	SpecializationName string
	Invitation         string
	AmountEmployee     int64
	AmountCourse       int64
	Courses            []course.Domain
	Users              []users.Domain
}

func NewSpecialization(id, companyId, name, invitation string) Domain {
	return Domain{
		ID:         id,
		CompanyID:  companyId,
		Name:       name,
		Invitation: invitation,
	}
}

func (old *Domain) ModifySpecialization(name string) Domain {
	return Domain{
		ID:         old.ID,
		CompanyID:  old.CompanyID,
		Name:       name,
		Invitation: old.Invitation,
		CreatedAt:  old.CreatedAt,
	}
}
