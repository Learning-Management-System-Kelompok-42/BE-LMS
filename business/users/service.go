package users

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Auth struct {
	Token  string
	UserID string
}

type Claims struct {
	Email       string
	UserID      string
	LevelAccess string
	jwt.StandardClaims
}

type UserRepository interface {
	// Insert creates a new user
	Insert(user User) (id string, err error)

	// Update updates an existing user
	Update(user User) (err error)

	// GetByID returns a user by ID
	GetByID(id string) (user *User, err error)

	// Login LoginUser logs in a user
	Login(email, password string) (user User, err error)

	// CheckEmail checks if an email is already registered
	CheckEmail(email string) error
}

type UserService interface {
	// Register creates a new user
	Register(upsertUserSpec spec.UpsertUsersSpec) (id string, err error)

	// UpdateUser updates an existing user
	UpdateUser(user User, id string) (err error)

	// GetUserByID returns a user by ID
	GetUserByID(id string) (*User, error)

	// LoginUser logs in a user
	LoginUser(upsertLoginSpec spec.UpsertLoginUsersSpec) (Auth, error)
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

	newId := uuid.New().String()

	passwordHash := encrypt.HashPassword(upsertUserSpec.Password)

	newUser := NewUser(
		newId,
		upsertUserSpec.FullName,
		upsertUserSpec.Email,
		passwordHash,
		upsertUserSpec.Phone,
		upsertUserSpec.Address,
		upsertUserSpec.Role,
		upsertUserSpec.CompanyID,
		upsertUserSpec.LevelAccess,
	)

	email := upsertUserSpec.Email

	err = s.userRepo.CheckEmail(email)
	if err != nil {
		if err == exception.ErrEmailExists {
			return "", exception.ErrEmailExists
		}
	}

	id, err = s.userRepo.Insert(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *userService) UpdateUser(user User, id string) (err error) {
	return err
}

func (s *userService) GetUserByID(id string) (*User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return nil, exception.ErrDataNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return user, nil
}

func (s *userService) LoginUser(upsertLoginSpec spec.UpsertLoginUsersSpec) (Auth, error) {
	var err error
	return Auth{}, err
}
