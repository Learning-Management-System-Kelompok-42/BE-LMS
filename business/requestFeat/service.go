package requestFeat

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RequestFeatRepository interface {
	// InsertRequestFeat insert new request feat
	InsertRequestFeat(domain Domain) (id string, err error)
}

type RequestFeatService interface {
	// CreateRequestFeat create new request feat
	CreateRequestFeat(upsertRequestFeat spec.UpsertRequestFeat) (id string, err error)
}

type requestFeatService struct {
	repo     RequestFeatRepository
	validate *validator.Validate
}

func NewRequestFeatService(repo RequestFeatRepository) RequestFeatService {
	return &requestFeatService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *requestFeatService) CreateRequestFeat(upsertRequestFeat spec.UpsertRequestFeat) (id string, err error) {
	err = s.validate.Struct(upsertRequestFeat)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	newID := uuid.New().String()

	if upsertRequestFeat.RequestTime != "" || upsertRequestFeat.RequestDate != "" {
		newRequest := NewRequestFeatCounseling(
			newID,
			upsertRequestFeat.CompanyID,
			upsertRequestFeat.RequestType,
			upsertRequestFeat.Title,
			upsertRequestFeat.Reason,
			upsertRequestFeat.RequestDate,
			upsertRequestFeat.RequestTime,
		)

		id, err = s.repo.InsertRequestFeat(newRequest)
		if err != nil {
			return "", exception.ErrInternalServer
		}

		return id, nil
	}

	newRequest := NewRequestFeat(
		newID,
		upsertRequestFeat.CompanyID,
		upsertRequestFeat.RequestType,
		upsertRequestFeat.Title,
		upsertRequestFeat.Reason,
	)

	id, err = s.repo.InsertRequestFeat(newRequest)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
