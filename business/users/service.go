package users

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/encrypt"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserRepository interface {
	// Insert creates a new user
	Insert(user Domain) (id string, err error)

	// UpdateSpecializationName Update updates an existing user
	UpdateSpecializationName(userUpdate Domain) (id string, err error)

	// UpdateProfile updates an existing user
	UpdateProfile(userUpdate Domain) (id string, err error)

	// UpdatePassword updates an existing user
	UpdatePassword(userUpdate Domain) (id string, err error)

	// FindByID returns a user by ID
	FindByID(id string) (user Domain, err error)

	// FindAllUsers GetAllUsers returns all users
	FindAllUsers(companyID string) (users []Domain, err error)

	// FindAllUserByCourseID returns all users by course ID
	FindAllUserByCourseID(courseID string) (users []Domain, err error)

	// FindDetailUserDashboard returns a user by ID
	FindDetailUserDashboard(userID string) (user UserDetailDashboard, err error)

	// FindDetailCourseDashboardUsers returns a course by ID, this course will be return if user already enroll on courses
	FindDetailCourseDashboardUsers(userID string) (courses []CourseDetailDashboardUser, err error)

	// FindAllUserBySpecializationID returns all users by specialization ID
	FindAllUserBySpecializationID(specializationID string) (users []Domain, err error)

	// FindUserDashboard return detail employee for dashboard
	FindUserDashboard(employeeID string) (user DetailEmployeeDashboard, err error)

	// FindAllCourseByEmployeeID return top progress course
	FindAllCourseByEmployeeID(employeeID string) (domain []TopCourseProgress, err error)

	// FindLastOpenCourseByEmployeeID return last open course
	FindLastOpenCourseByEmployeeID(employeeID string) (domain []LastCourseOpen, err error)

	// CountAllCourseByUserID return amount of all course by employee id
	CountAllCourseByUserID(employeeID string) (count int64, err error)

	// CountCourseCompleted return amount of course completed
	CountCourseCompleted(employeeID string) (count int64, err error)

	// CountCourseIncomplete return amount of course incomplete
	CountCourseIncomplete(employeeID string) (count int64, err error)

	// CountModulesByCourseID return amount of modules
	CountModulesByCourseID(courseID string) (count int64, err error)

	// CountModulesCompletedByEmployeeID return amount of modules completed
	CountModulesCompletedByEmployeeID(courseID, employeeID string) (count int64, err error)

	// CheckEmail checks if an email is already registered
	CheckEmail(email string) error
}

type UserService interface {
	// Register creates a new user
	Register(upsertUserSpec spec.UpsertUsersSpec) (id string, err error)

	// UpdateSpecializationName UpdateUser updates an existing user
	UpdateSpecializationName(upsertUpdateSpecName spec.UpsertUpdateSpecName) (id string, err error)

	// UpdateProfile updates an existing user
	UpdateProfile(upsertUpdateProfile spec.UpsertUpdateProfileSpec) (id string, err error)

	// UpdatePassword updates an existing user
	UpdatePassword(upsertUpdatePassowrd spec.UpsertUpdatePassword) (id string, err error)

	// GetDetailUserByID returns a user by ID
	GetDetailUserByID(id string) (*Domain, error)

	// GetDetailUserDashboard  return a user by id
	GetDetailUserDashboard(userID string) (user ToResponseDetailUserDashboard, err error)

	//GetDashboardEmployee return data for dashboard employee
	GetDashboardEmployee(employeeID, specializationID string) (domain DashboardEmployee, err error)

	// GetAllUsers returns all users
	GetAllUsers(companyID string) (users []Domain, err error)
}

type userService struct {
	userRepo UserRepository
	// enrollmentRepo enrollments.EnrollmentRepository
	validate *validator.Validate
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
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
		if err == exception.ErrEmployeeNotFound {
			return result, exception.ErrEmployeeNotFound
		}

		return result, exception.ErrInternalServer
	}

	course, err := s.userRepo.FindDetailCourseDashboardUsers(user.ID)
	if err != nil {
		// if err == exception.ErrNotFound {
		// 	return result, exception.ErrNotFound
		// }

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

func (s *userService) GetDashboardEmployee(employeeID, specializationID string) (domain DashboardEmployee, err error) {
	// Get user detail
	user, err := s.userRepo.FindUserDashboard(employeeID)
	if err != nil {
		if err == exception.ErrEmployeeNotFound {
			return domain, exception.ErrEmployeeNotFound
		}

		return domain, exception.ErrInternalServer
	}

	// Amount of Course
	amountCourse, err := s.userRepo.CountAllCourseByUserID(employeeID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	// kursus selesai
	amountCourseCompleted, err := s.userRepo.CountCourseCompleted(employeeID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	// kursus belum selesai
	// Don't forget to change business logic when user give feedback
	// Change if he give feedback, then update status into true
	amountCourseIncomplete, err := s.userRepo.CountCourseIncomplete(employeeID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	domain = DashboardEmployee{
		DetailEmployee:         user,
		AmountCourse:           amountCourse,
		AmountCourseCompleted:  amountCourseCompleted,
		AmountCourseIncomplete: amountCourseIncomplete,
	}

	// Get all course
	courses, err := s.userRepo.FindAllCourseByEmployeeID(employeeID)
	if err != nil {
		return domain, exception.ErrInternalServer
	}

	// Get top 4 highest progress course
	for _, v := range courses {
		// Calculate percentage progress course by user
		totalModule, _ := s.userRepo.CountModulesByCourseID(v.CourseID)
		moduleCompleted, _ := s.userRepo.CountModulesCompletedByEmployeeID(v.CourseID, employeeID)
		// percentage total module completed from course
		var percentageModule int64
		if totalModule != 0 {
			percentageModule = (moduleCompleted * 100) / totalModule
		}

		fmt.Println("percentageModule ", percentageModule)

		progress := TopCourseProgress{
			CourseID:  v.CourseID,
			Title:     v.Title,
			Thumbnail: v.Thumbnail,
			Progress:  percentageModule,
		}

		domain.TopCourseProgress = append(domain.TopCourseProgress, progress)
	}

	// kursus sering di buka 7 hari
	coursesLatestEnroll, err := s.userRepo.FindLastOpenCourseByEmployeeID(employeeID)
	if err != nil {
		if err == exception.ErrCourseNotFound {
			return domain, exception.ErrCourseNotFound
		}

		return domain, exception.ErrInternalServer
	}

	domain.TopCourseOften7Days = coursesLatestEnroll

	return domain, nil
}
