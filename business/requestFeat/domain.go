package requestFeat

import "time"

type Domain struct {
	ID          string
	CompanyID   string
	RequestType string
	Title       string
	Reason      string
	Status      bool
	RequestDate string
	RequestTime string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRequestFeat(id, companyID, requestType, title, reason string) Domain {
	return Domain{
		ID:          id,
		CompanyID:   companyID,
		RequestType: requestType,
		Title:       title,
		Reason:      reason,
		Status:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}

func NewRequestFeatCounseling(id, companyID, requestType, title, reason, requestDate, requestTime string) Domain {
	return Domain{
		ID:          id,
		CompanyID:   companyID,
		RequestType: requestType,
		Title:       title,
		Reason:      reason,
		Status:      false,
		RequestDate: requestDate,
		RequestTime: requestTime,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
