package company

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CompanyRepository interface {
	// Insert creates a new company into database
	Insert(company Company) (id string, err error)

	// CheckEmail checks if an email is already registered
	CheckEmail(email string) error
}

type CompanyService interface {
	// Register creates a new company
	Register(upsertCompanySpec spec.UpsertCompanySpec) (id string, err error)
}

type companyService struct {
	companyRepository CompanyRepository
	validate          *validator.Validate
}

func NewCompanyService(repo CompanyRepository) CompanyService {
	return &companyService{
		companyRepository: repo,
		validate:          validator.New(),
	}
}

func (s *companyService) Register(upsertCompanySpec spec.UpsertCompanySpec) (id string, err error) {
	err = s.validate.Struct(&upsertCompanySpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	fmt.Println("error controller = ", &upsertCompanySpec)

	newId := uuid.New().String()

	newCompany := NewCompany(
		newId,
		upsertCompanySpec.Name,
		upsertCompanySpec.Address,
		upsertCompanySpec.Web,
		upsertCompanySpec.Email,
		upsertCompanySpec.Sector,
		upsertCompanySpec.Logo,
	)

	fmt.Println("new company = ", newCompany)

	email := upsertCompanySpec.Email

	if err := s.companyRepository.CheckEmail(email); err != nil {
		if err == exception.ErrEmailExists {
			return "", exception.ErrEmailExists
		}

		return "", exception.ErrInternalServer
	}

	id, err = s.companyRepository.Insert(newCompany)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
