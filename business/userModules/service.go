package userModules

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserModulesRepository interface {
	// InsertProgress insert a new progress
	InsertProgress(progress Domain) (id string, err error)

	// UpdateProgress insert a new progress
	UpdateProgress(progress Domain) (id string, err error)

	// FindProgressByModuleID find a progress by module ID
	FindProgressByModuleID(moduleID string) (progress *Domain, err error)

	// CheckProgressExist check if progress exist
	CheckProgressExist(userID, courseID, moduleID string) (err error)
}

type UserModulesService interface {
	// Create insert a new progress
	CreateProgress(upsertProgressSpec spec.UpsertProgressSpec) (id string, err error)

	// UpdateProgress update a new progress modules
	UpdateProgress(upsertProgressSpec spec.UpsertProgressSpec) (id string, err error)
}

type userModulesService struct {
	userModulesRepo UserModulesRepository
	validate        *validator.Validate
}

func NewUserModulesService(userModulesRepo UserModulesRepository) UserModulesService {
	return &userModulesService{
		userModulesRepo: userModulesRepo,
		validate:        validator.New(),
	}
}

func (s *userModulesService) UpdateProgress(upsertProgressSpec spec.UpsertProgressSpec) (id string, err error) {
	err = s.validate.Struct(&upsertProgressSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldProgress, err := s.userModulesRepo.FindProgressByModuleID(upsertProgressSpec.ModuleID)
	if err != nil {
		if err == exception.ErrModuleNotFound {
			return "", exception.ErrModuleNotFound
		}

		return "", exception.ErrInternalServer
	}

	newProgress := oldProgress.UpdateProgress(
		upsertProgressSpec.Point,
		upsertProgressSpec.Status,
	)

	id, err = s.userModulesRepo.UpdateProgress(newProgress)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *userModulesService) CreateProgress(upsertProgressSpec spec.UpsertProgressSpec) (id string, err error) {
	err = s.validate.Struct(&upsertProgressSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	err = s.userModulesRepo.CheckProgressExist(
		upsertProgressSpec.UserID,
		upsertProgressSpec.CourseID,
		upsertProgressSpec.ModuleID,
	)
	if err != nil {
		return "", exception.ErrProgressAlreadyExist
	}

	newID := uuid.New().String()
	newProgress := NewProggresCourse(
		newID,
		upsertProgressSpec.UserID,
		upsertProgressSpec.CourseID,
		upsertProgressSpec.ModuleID,
		upsertProgressSpec.Point,
		upsertProgressSpec.Status,
	)

	id, err = s.userModulesRepo.InsertProgress(newProgress)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
