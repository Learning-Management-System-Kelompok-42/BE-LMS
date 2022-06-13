package company

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	cloud "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/cloudinary"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CompanyRepository interface {
	// Insert creates a new company into database
	Insert(company Domain) (id string, err error)

	// CheckEmail checks if an email is already registered
	CheckWeb(web string) error
}

type CompanyService interface {
	// Register creates a new company
	Register(upsertCompanySpec spec.UpsertCompanySpec) (id string, err error)
}

type companyService struct {
	companyRepo CompanyRepository
	userRepo    users.UserRepository
	validate    *validator.Validate
}

func NewCompanyService(companyRepo CompanyRepository, userRepo users.UserRepository) CompanyService {
	return &companyService{
		companyRepo: companyRepo,
		userRepo:    userRepo,
		validate:    validator.New(),
	}
}

func (s *companyService) Register(upsertCompanySpec spec.UpsertCompanySpec) (id string, err error) {
	// Next we will upgrade query into database transactional
	err = s.validate.Struct(&upsertCompanySpec)
	if err != nil {
		return "", err
	}

	err = s.companyRepo.CheckWeb(upsertCompanySpec.Website)
	if err != nil {
		return "", exception.ErrWebExists
	}

	err = s.userRepo.CheckEmail(upsertCompanySpec.EmailAdmin)
	if err != nil {
		return "", exception.ErrEmailExists
	}

	urlLogo, err := cloud.ImageUploadHelper(upsertCompanySpec.Logo)
	if err != nil {
		return "", exception.ErrCantUploadImage
	}
	fmt.Print("urlLogo = ", urlLogo)

	idCompany := uuid.New().String()
	newCompany := NewCompany(
		idCompany,
		upsertCompanySpec.NameCompany,
		upsertCompanySpec.AddressCompany,
		upsertCompanySpec.Website,
		upsertCompanySpec.Sector,
		urlLogo,
	)

	id, err = s.companyRepo.Insert(newCompany)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	idUser := uuid.New().String()
	role := "admin"
	specializationID := "SPEC-002"
	hashPassword := encrypt.HashPassword(upsertCompanySpec.PasswordAdmin)

	newUser := users.NewUser(
		idUser,
		idCompany,
		specializationID,
		role,
		upsertCompanySpec.NameAdmin,
		upsertCompanySpec.EmailAdmin,
		hashPassword,
		upsertCompanySpec.PhoneNumber,
		upsertCompanySpec.AddressAdmin,
		role,
	)

	_, err = s.userRepo.Insert(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
