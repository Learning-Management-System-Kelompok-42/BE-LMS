package requestFeat

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat"
	"gorm.io/gorm"
)

type RequestCourse struct {
	ID          string `gorm:"primaryKey,size:200"`
	CompanyID   string `gorm:"size:200"`
	RequestType string
	Title       string
	Reason      string
	Status      bool
	RequestDate string
	RequestTime string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain requestFeat.Domain) RequestCourse {
	return RequestCourse{
		ID:          domain.ID,
		CompanyID:   domain.CompanyID,
		RequestType: domain.RequestType,
		Title:       domain.Title,
		Reason:      domain.Reason,
		Status:      domain.Status,
		RequestDate: domain.RequestDate,
		RequestTime: domain.RequestTime,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   gorm.DeletedAt{},
	}
}

func (request *RequestCourse) ToDomain() requestFeat.Domain {
	return requestFeat.Domain{
		ID:          request.ID,
		CompanyID:   request.CompanyID,
		RequestType: request.RequestType,
		Title:       request.Title,
		Reason:      request.Reason,
		Status:      request.Status,
		RequestDate: request.RequestDate,
		RequestTime: request.RequestTime,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}
}
