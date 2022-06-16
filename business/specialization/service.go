package specialization

import (
	"fmt"
	"strings"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SpecializationRepository interface {
	// Insert creates a new specialization into database
	Insert(specialization Domain) (id string, err error)
}

type SpecializationService interface {
	// Register creates a new specialization
	Register(upsertSpecializationSpec spec.UpsertSpecializationSpec) (id string, err error)
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
	fmt.Println("separateLink", separateLink)

	newId := uuid.New().String()

	newSpec := NewSpecialization(
		newId,
		upsertSpecializationSpec.Name,
		separateLink,
	)

	id, err = s.specializationRepo.Insert(newSpec)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
