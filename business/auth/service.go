package auth

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/repository/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

type AuthRepository interface {
	// Login find user by email and password
	Login(email string) (user users.User, err error)
}

type AuthService interface {
	// Login
	LoginUser(upsertAuth spec.UpsertAuthSpec) (auth *Auth, err error)
}

type authService struct {
	authRepo AuthRepository
	validate *validator.Validate
	cfg      *config.AppConfig
}

func NewAuthService(authRepo AuthRepository, config *config.AppConfig) AuthService {
	return &authService{
		authRepo: authRepo,
		validate: validator.New(),
		cfg:      config,
	}
}

func (s *authService) LoginUser(upsertAuth spec.UpsertAuthSpec) (auth *Auth, err error) {
	err = s.validate.Struct(&upsertAuth)
	if err != nil {
		return nil, exception.ErrInvalidRequest
	}

	user, err := s.authRepo.Login(upsertAuth.Email)
	if err != nil {
		if err == exception.ErrNotFound {
			return nil, exception.ErrNotFound
		}

		return nil, exception.ErrInternalServer
	}

	hashed := encrypt.CheckPasswordHash(upsertAuth.Password, user.Password)
	if !hashed {
		return nil, exception.ErrWrongPassword
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		CompanyID:        user.CompanyID,
		SpecializationID: user.SpecializationID,
		UserID:           user.ID,
		LevelAccess:      user.LevelAccess,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.cfg.App.SecretKey))

	if err != nil {
		return nil, exception.ErrInvalidGenerateToken
	}

	auth = &Auth{
		Token:            tokenString,
		UserID:           user.ID,
		CompanyID:        user.CompanyID,
		LevelAccess:      user.LevelAccess,
		SpecializationID: user.SpecializationID,
	}

	return auth, nil
}
