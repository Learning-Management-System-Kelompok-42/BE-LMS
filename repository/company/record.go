package company

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/requestCourse"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"gorm.io/gorm"
)

type Company struct {
	ID             string `gorm:"primaryKey,size:200"`
	Name           string
	Address        string
	Web            string `gorm:"size:250;uniqueIndex"`
	Sector         string
	Logo           string
	Courses        []course.Course                 `gorm:"foreignkey:CompanyID"`
	Users          []users.User                    `gorm:"foreignKey:CompanyID"`
	RequestCourses []requestCourse.RequestCourse   `gorm:"foreignKey:CompanyID"`
	Specialization []specialization.Specialization `gorm:"foreignKey:CompanyID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (comp *Company) ToDomain() *company.Domain {
	return &company.Domain{
		ID:        comp.ID,
		Name:      comp.Name,
		Address:   comp.Address,
		Web:       comp.Web,
		Sector:    comp.Sector,
		Logo:      comp.Logo,
		CreatedAt: comp.CreatedAt,
		UpdatedAt: comp.UpdatedAt,
	}
}

func FromDomain(domain company.Domain) Company {
	return Company{
		ID:             domain.ID,
		Name:           domain.Name,
		Address:        domain.Address,
		Web:            domain.Web,
		Sector:         domain.Sector,
		Logo:           domain.Logo,
		Users:          nil,
		RequestCourses: nil,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      gorm.DeletedAt{},
	}
}
