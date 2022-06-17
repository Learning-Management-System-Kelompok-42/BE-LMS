package users

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserRepository interface {
	// Insert creates a new user
	Insert(user Domain) (id string, err error)

	// Update updates an existing user
	Update(user Domain) (err error)

	// FindByID returns a user by ID
	FindByID(id string) (user Domain, err error)

	// GetAllUsers returns all users
	FindAllUsers(userID string) (users []Domain, err error)

	// CheckEmail checks if an email is already registered
	CheckEmail(email string) error
}

type UserService interface {
	// Register creates a new user
	Register(upsertUserSpec spec.UpsertUsersSpec) (id string, err error)

	// UpdateUser updates an existing user
	UpdateUser(user Domain, id string) (err error)

	// GetUserByID returns a user by ID
	GetUserByID(id string) (*Domain, error)

	// GetAllUsers returns all users
	GetAllUsers(userID string) (users []Domain, err error)
}

type userService struct {
	userRepo UserRepository
	validate *validator.Validate
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		userRepo: repo,
		validate: validator.New(),
	}
}

func (s *userService) Register(upsertUserSpec spec.UpsertUsersSpec) (id string, err error) {
	err = s.validate.Struct(&upsertUserSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	err = s.userRepo.CheckEmail(upsertUserSpec.Email)
	if err != nil {
		if err == exception.ErrEmailExists {
			return "", exception.ErrEmailExists
		}
	}

	newId := uuid.New().String()
	passwordHash := encrypt.HashPassword(upsertUserSpec.Password)
	levelAccess := "employee"

	newUser := NewUser(
		newId,
		upsertUserSpec.CompanyID,
		upsertUserSpec.SpecializationID,
		upsertUserSpec.Role,
		upsertUserSpec.FullName,
		upsertUserSpec.Email,
		passwordHash,
		upsertUserSpec.Phone,
		upsertUserSpec.Address,
		levelAccess,
	)

	id, err = s.userRepo.Insert(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *userService) UpdateUser(user Domain, id string) (err error) {
	return err
}

func (s *userService) GetUserByID(id string) (*Domain, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return nil, exception.ErrDataNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return &user, nil
}

func (s *userService) GetAllUsers(userID string) (users []Domain, err error) {
	users, err = s.userRepo.FindAllUsers(userID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return nil, exception.ErrDataNotFound
		}

		return nil, exception.ErrInternalServer
	}

	return users, nil
}
