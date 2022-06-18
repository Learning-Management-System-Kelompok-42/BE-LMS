package module

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ModuleRepository interface {
	// Insert insert a new module
	Insert(domain Domain) (id string, err error)
}

type ModuleService interface {
	// Insert insert a new module
	Create(UpsertModuleSpec spec.UpsertModuleSpec) (id string, err error)
}

type moduleService struct {
	repo     ModuleRepository
	validate *validator.Validate
}

func NewModuleService(repo ModuleRepository) ModuleService {
	return &moduleService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *moduleService) Create(UpsertModuleSpec spec.UpsertModuleSpec) (id string, err error) {
	err = s.validate.Struct(&UpsertModuleSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	moduleID := uuid.New().String()

	newModule := NewModule(
		moduleID,
		UpsertModuleSpec.CourseID,
		UpsertModuleSpec.YoutubeURL,
		UpsertModuleSpec.SlideURL,
		UpsertModuleSpec.Title,
		UpsertModuleSpec.Orders,
	)

	id, err = s.repo.Insert(newModule)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
