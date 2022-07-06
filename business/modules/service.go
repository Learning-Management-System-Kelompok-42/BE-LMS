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

	// GetByID get a module by id
	FindByID(id string) (domain Domain, err error)

	// Update update a module
	Update(domain Domain) (id string, err error)
}

type ModuleService interface {
	// Insert insert a new module
	Create(UpsertModuleSpec spec.UpsertModuleSpec) (id string, err error)

	// GetByID get a module by id
	GetByID(id string) (domain Domain, err error)

	// Update update a module
	Update(UpsertModuleSpec spec.UpsertModuleSpec) (id string, err error)
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
		if err == exception.ErrCourseNotFound {
			return id, exception.ErrCourseNotFound
		}
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *moduleService) Update(UpsertModuleSpec spec.UpsertModuleSpec) (id string, err error) {
	err = s.validate.Struct(&UpsertModuleSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldModule, err := s.repo.FindByID(UpsertModuleSpec.ID)
	if err != nil {
		if err == exception.ErrNotFound {
			return id, exception.ErrNotFound
		}

		return id, exception.ErrInternalServer
	}

	newModule := oldModule.ModifyModule(
		UpsertModuleSpec.Title,
		UpsertModuleSpec.YoutubeURL,
		UpsertModuleSpec.SlideURL,
		UpsertModuleSpec.Orders,
	)

	id, err = s.repo.Update(newModule)
	if err != nil {
		if err == exception.ErrNotFound {
			return id, exception.ErrNotFound
		}

		return id, exception.ErrInternalServer
	}

	return id, nil
}

func (s *moduleService) GetByID(id string) (domain Domain, err error) {
	domain, err = s.repo.FindByID(id)

	if err != nil {
		if err == exception.ErrNotFound {
			return domain, exception.ErrNotFound
		}

		return domain, exception.ErrInternalServer
	}

	return domain, nil
}
