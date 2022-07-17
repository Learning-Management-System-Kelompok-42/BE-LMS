package company

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/s3"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CompanyRepository interface {
	// Insert creates a new company into database
	Insert(company Domain) (id string, err error)

	// CountSpecializationByCompanyID return amount of specialization
	CountSpecializationByCompanyID(companyID string) (count int64, err error)

	// CountEmployeeByCompanyID return amount of employee
	CountEmployeeByCompanyID(companyID string) (count int64, err error)

	// CountCourseByCompanyID return amount of course
	CountCourseByCompanyID(companyID string) (count int64, err error)

	// FindSpecializationEmployee count total specialization with the highest amount of employee
	FindSpecializationEmployee(companyID string) (domain []DashboardCompanySpecialization, err error)

	// FindCourseEmployee count total course with highest amount of employee
	FindCourseEmployee(companyID string) (domain []DashboardCompanyCourse, err error)

	// FindCompanyByCompanyID return company by companyID
	FindCompanyByCompanyID(companyID string) (domain DashboardCompanyAdmin, err error)

	// NewRequest return new request

	// FindCompanyByID return company by id
	FindCompanyByID(companyID string) (domain *Domain, err error)

	// UpdateProfile update profile company
	UpdateProfile(company Domain) (id string, err error)

	// CheckWeb checks if an email is already registered
	CheckWeb(web string) error
}

type CompanyService interface {
	// Register creates a new company
	Register(upsertCompanySpec spec.UpsertCompanySpec) (id string, err error)

	// Dashboard returns a list of specialization and amount of employees
	Dashboard(companyID string) (domain DashboardDomain, err error)

	// UpdateProfile updates a company profile
	UpdateProfile(upsertProfileCompanySpec spec.UpsertProfileCompanySpec) (id string, err error)

	// GetCompanyByID returns a company profile
	GetCompanyByID(companyID string) (domain Domain, err error)
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

	// urlLogo, err := cloud.ImageUploadHelper(upsertCompanySpec.Logo)
	// if err != nil {
	// 	return "", exception.ErrCantUploadImage
	// }

	urlLogo, err := s3.UploadFileHelper(upsertCompanySpec.Logo, upsertCompanySpec.FileName)
	if err != nil {
		return "", exception.ErrCantUploadImage
	}

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
	specializationID := "SPEC-001"
	hashPassword := encrypt.HashPassword(upsertCompanySpec.PasswordAdmin)

	newUser := users.NewUser(
		idUser,
		idCompany,
		specializationID,
		upsertCompanySpec.NameAdmin,
		upsertCompanySpec.EmailAdmin,
		hashPassword,
		upsertCompanySpec.PhoneNumber,
		upsertCompanySpec.AddressAdmin,
		role,
		role,
	)

	_, err = s.userRepo.Insert(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *companyService) Dashboard(companyID string) (domain DashboardDomain, err error) {
	company, err := s.companyRepo.FindCompanyByCompanyID(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	amountSpecialization, err := s.companyRepo.CountSpecializationByCompanyID(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	amountEmployee, err := s.companyRepo.CountEmployeeByCompanyID(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	amountCourse, err := s.companyRepo.CountCourseByCompanyID(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	specialization, err := s.companyRepo.FindSpecializationEmployee(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	course, err := s.companyRepo.FindCourseEmployee(companyID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	domain = DashboardDomain{
		UserID:               company.UserID,
		CompanyID:            company.CompanyID,
		NameAdmin:            company.NameAdmin,
		NameCompany:          company.NameCompany,
		AmountSpecialization: amountSpecialization,
		AmountEmployee:       amountEmployee,
		AmountCourse:         amountCourse,
		Specialization:       specialization,
		Course:               course,
	}

	return domain, nil
}

func (s *companyService) GetCompanyByID(companyID string) (domain Domain, err error) {
	company, err := s.companyRepo.FindCompanyByID(companyID)
	if err != nil {
		if err == exception.ErrCompanyNotFound {
			return domain, exception.ErrCompanyNotFound
		}

		return domain, exception.ErrInternalServer
	}

	domain = *company

	return domain, nil
}

func (s *companyService) UpdateProfile(upsertProfileCompanySpec spec.UpsertProfileCompanySpec) (id string, err error) {
	err = s.validate.Struct(&upsertProfileCompanySpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldCompany, err := s.companyRepo.FindCompanyByID(upsertProfileCompanySpec.CompanyID)
	if err != nil {
		if err == exception.ErrCompanyNotFound {
			return "", exception.ErrCompanyNotFound
		}

		return "", exception.ErrInternalServer
	}

	imageURL, err := s3.UploadFileHelper(upsertProfileCompanySpec.Logo, upsertProfileCompanySpec.FileName)
	if err != nil {
		return "", exception.ErrCantUploadImage
	}

	newCompany := oldCompany.ModifyCompany(
		upsertProfileCompanySpec.NameCompany,
		upsertProfileCompanySpec.AddressCompany,
		upsertProfileCompanySpec.Website,
		upsertProfileCompanySpec.Sector,
		imageURL,
	)

	id, err = s.companyRepo.UpdateProfile(newCompany)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
