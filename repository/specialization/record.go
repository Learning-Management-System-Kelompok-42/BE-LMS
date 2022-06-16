package specialization

import (
	"time"

	specializations "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specializationCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"gorm.io/gorm"
)

type Specialization struct {
	ID                    string `gorm:"primaryKey;size:200"`
	Name                  string
	Invitation            string
	Users                 []users.User                                `gorm:"foreignKey:SpecializationID"`
	SpecializationCourses []specializationCourse.SpecializationCourse `gorm:"foreignKey:SpecializationID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

func (specialization *Specialization) ToDomain() specializations.Domain {
	return specializations.Domain{
		ID:         specialization.ID,
		Name:       specialization.Name,
		Invitation: specialization.Invitation,
	}
}

func ToDomainList(specialization []Specialization) []specializations.Domain {
	var domains []specializations.Domain
	for _, spec := range specialization {
		domains = append(domains, spec.ToDomain())
	}
	return domains
}

func FromDomain(domain specializations.Domain) Specialization {
	return Specialization{
		ID:                    domain.ID,
		Name:                  domain.Name,
		Invitation:            domain.Invitation,
		Users:                 nil,
		SpecializationCourses: nil,
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
		DeletedAt:             gorm.DeletedAt{},
	}
}
