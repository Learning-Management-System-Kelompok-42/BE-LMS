package specialization

import (
	"strings"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SpecializationRepository interface {
	// Insert creates a new specialization into database
	Insert(specialization Domain) (id string, err error)

	// FindInvitation
	FindInvitation(invitation string) (specialization Domain, err error)

	// FindDashboardSpecialization returns all specialization
	FindDashboardSpecialization(companyID string) (specializations []Domain, err error)

	// CountCourse returns the number of course
	CountCourse(specID string) (result int64, err error)

	// CountEmployee returns the number of employee
	CountEmployee(companyID, specID string) (result int64, err error)
}

type SpecializationService interface {
	// Register creates a new specialization
	Register(upsertSpecializationSpec spec.UpsertSpecializationSpec) (id string, err error)

	// GetInvitation returns a specialization by invitation
	GetInvitation(invitation string) (specialization Domain, err error)

	// GetAllSpecialization returns all specialization
	GetAllSpecialization(companyID string) (specializations []SpecializationDashboard, err error)
}

type specializationService struct {
	specializationRepo SpecializationRepository
	validate           *validator.Validate
}

func NewSpecializationService(specializationRepo SpecializationRepository) SpecializationService {
	return &specializationService{
		specializationRepo: specializationRepo,
		validate:           validator.New(),
	}
}

func (s *specializationService) Register(upsertSpecializationSpec spec.UpsertSpecializationSpec) (id string, err error) {
	err = s.validate.Struct(&upsertSpecializationSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	link := upsertSpecializationSpec.Invitation
	separateLink := strings.SplitAfter(link, "invitation/")[1]

	newId := uuid.New().String()

	newSpec := NewSpecialization(
		newId,
		upsertSpecializationSpec.CompanyID,
		upsertSpecializationSpec.Name,
		separateLink,
	)

	id, err = s.specializationRepo.Insert(newSpec)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *specializationService) GetInvitation(invitation string) (specialization Domain, err error) {
	link := strings.SplitAfter(invitation, "invitation/")[1]

	specialization, err = s.specializationRepo.FindInvitation(link)
	if err != nil {
		if err == exception.ErrNotFound {
			return specialization, exception.ErrNotFound
		}
		return specialization, exception.ErrInternalServer
	}

	return specialization, nil
}

func (s *specializationService) GetAllSpecialization(companyID string) (specializations []SpecializationDashboard, err error) {
	// Get all specialization
	specialization, err := s.specializationRepo.FindDashboardSpecialization(companyID)
	if err != nil {
		return specializations, err
	}

	// count course
	for _, spec := range specialization {
		countCourse, err := s.specializationRepo.CountCourse(spec.ID)
		if err != nil {
			return specializations, err
		}

		countEmployee, err := s.specializationRepo.CountEmployee(companyID, spec.ID)

		specializations = append(specializations, SpecializationDashboard{
			SpecializationID:   spec.ID,
			SpecializationName: spec.Name,
			AmountEmployee:     countEmployee,
			AmountCourse:       countCourse,
		})
	}

	return specializations, nil
}
