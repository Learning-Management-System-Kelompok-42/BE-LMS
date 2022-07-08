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
	UpdateSpecializationName(userUpdate Domain) (id string, err error)

	// UpdateProfile updates an existing user
	UpdateProfile(userUpdate Domain) (id string, err error)

	// UpdatePassword updates an existing user
	UpdatePassword(userUpdate Domain) (id string, err error)

	// FindByID returns a user by ID
	FindByID(id string) (user Domain, err error)

	// GetAllUsers returns all users
	FindAllUsers(companyID string) (users []Domain, err error)

	// FindAllUserByCourseID returns all users by course ID
	FindAllUserByCourseID(courseID string) (users []Domain, err error)

	// FindDetailUserDashboard returns a user by ID
	FindDetailUserDashboard(userID string) (user UserDetailDashboard, err error)

	// FindDetailCourseDashboardUsers returns a course by ID, this course will be return if user already enroll on courses
	FindDetailCourseDashboardUsers(userID string) (courses []CourseDetailDashboardUser, err error)

	// FindAllUserBySpecializationID returns all users by specialization ID
	FindAllUserBySpecializationID(specializationID string) (users []Domain, err error)

	// CheckEmail checks if an email is already registered
	CheckEmail(email string) error
}

type UserService interface {
	// Register creates a new user
	Register(upsertUserSpec spec.UpsertUsersSpec) (id string, err error)

	// UpdateUser updates an existing user
	UpdateSpecializationName(upsertUpdateSpecName spec.UpsertUpdateSpecName) (id string, err error)

	// UpdateProfile updates an existing user
	UpdateProfile(upsertUpdateProfile spec.UpsertUpdateProfileSpec) (id string, err error)

	// UpdatePassword updates an existing user
	UpdatePassword(upsertUpdatePassowrd spec.UpsertUpdatePassword) (id string, err error)

	// GetUserByID returns a user by ID
	GetDetailUserByID(id string) (*Domain, error)

	// GetDetailUserDashboard  return a user by id
	GetDetailUserDashboard(userID string) (user ToResponseDetailUserDashboard, err error)

	// GetAllUsers returns all users
	GetAllUsers(companyID string) (users []Domain, err error)
}

type userService struct {
	userRepo UserRepository
	// enrollmentRepo enrollments.EnrollmentRepository
	validate *validator.Validate
}

func NewUserService(user UserRepository) UserService {
	return &userService{
		userRepo: user,
		// enrollmentRepo: enrollmentRepo,
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

func (s *userService) UpdateSpecializationName(upsertUpdateSpecName spec.UpsertUpdateSpecName) (id string, err error) {
	oldUser, err := s.userRepo.FindByID(upsertUpdateSpecName.UserID)
	if err != nil {
		if err == exception.ErrEmployeeNotFound {
			return "", exception.ErrEmployeeNotFound
		}
		return "", exception.ErrInternalServer
	}

	newUser := oldUser.ModifySpecializationName(upsertUpdateSpecName.SpecializationID)

	id, err = s.userRepo.UpdateSpecializationName(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *userService) GetDetailUserByID(id string) (*Domain, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return nil, exception.ErrDataNotFound
		}
		return nil, exception.ErrInternalServer
	}

	return &user, nil
}

func (s *userService) GetAllUsers(companyID string) (users []Domain, err error) {
	users, err = s.userRepo.FindAllUsers(companyID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return nil, exception.ErrDataNotFound
		}

		return nil, exception.ErrInternalServer
	}

	return users, nil
}

func (s *userService) GetDetailUserDashboard(userID string) (result ToResponseDetailUserDashboard, err error) {
	user, err := s.userRepo.FindDetailUserDashboard(userID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return result, exception.ErrDataNotFound
		}

		return result, exception.ErrInternalServer
	}

	course, err := s.userRepo.FindDetailCourseDashboardUsers(user.ID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return result, exception.ErrDataNotFound
		}

		return result, exception.ErrInternalServer
	}

	result = ToResponseDetailUserDashboard{
		User:    user,
		Courses: course,
	}

	return result, nil
}

func (s *userService) UpdateProfile(upsertUpdateProfile spec.UpsertUpdateProfileSpec) (id string, err error) {
	err = s.validate.Struct(&upsertUpdateProfile)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldUser, err := s.userRepo.FindByID(upsertUpdateProfile.ID)
	if err != nil {

		if err == exception.ErrEmployeeNotFound {
			return "", exception.ErrEmployeeNotFound
		}
		return "", exception.ErrInternalServer
	}

	newUser := oldUser.ModifyUser(
		upsertUpdateProfile.FullName,
		upsertUpdateProfile.Email,
		upsertUpdateProfile.PhoneNumber,
		upsertUpdateProfile.Address,
		oldUser.Avatar,
	)

	id, err = s.userRepo.UpdateProfile(newUser)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *userService) UpdatePassword(upsertUpdatePassowrd spec.UpsertUpdatePassword) (id string, err error) {
	err = s.validate.Struct(&upsertUpdatePassowrd)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldUser, err := s.userRepo.FindByID(upsertUpdatePassowrd.UserID)
	if err != nil {
		if err == exception.ErrEmployeeNotFound {
			return "", exception.ErrEmployeeNotFound
		}
		return "", exception.ErrInternalServer
	}

	if !encrypt.CheckPasswordHash(upsertUpdatePassowrd.OldPassword, oldUser.Password) {
		return "", exception.ErrWrongPassword
	}

	passwordHash := encrypt.HashPassword(upsertUpdatePassowrd.NewPassword)

	newPassword := oldUser.ModifyPassword(passwordHash)

	id, err = s.userRepo.UpdatePassword(newPassword)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
