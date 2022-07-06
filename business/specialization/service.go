package specialization

import (
	"strings"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SpecializationRepository interface {
	// Insert creates a new specialization into database
	Insert(specialization Domain) (id string, err error)

	// InsertCourseSpecialization creates a new course into specialization table
	InsertCourseSpecialization(courseID, specializationID string) (id string, err error)

	// FindInvitation
	FindInvitation(invitation string) (specialization Domain, err error)

	// FindDashboardSpecialization returns all specialization
	FindDashboardSpecialization(companyID string) (specializations []Domain, err error)

	// FindSpecializationByID returns specialization by id
	FindSpecializationByID(specializationID, companyID string) (specialization Domain, err error)

	// CountCourse returns the number of course
	CountCourse(specID string) (result int64, err error)

	// CountEmployee returns the number of employee
	CountEmployee(companyID, specID string) (result int64, err error)

	// CheckLinkInviation returns boolean
	CheckLinkInviation(link string) (err error)

	// UpdateSpecialization updates a specialization
	UpdateSpecialization(specialization Domain) (id string, err error)
}

type SpecializationService interface {
	// Register creates a new specialization
	Register(upsertSpecializationSpec spec.UpsertSpecializationSpec) (id string, err error)

	// AddCourseSpecialization adds a new course into specialization
	AddCourseSpecialization(upsertCourseSpecializationSpec spec.UpsertCourseSpecializationSpec) (id string, err error)

	// GetInvitation returns a specialization by invitation
	GetInvitation(invitation string) (specialization Domain, err error)

	// GetAllSpecialization returns all specialization
	GetAllSpecialization(companyID string) (specializations []SpecializationDashboard, err error)

	// GetSpecializationByID returns a specialization by id
	GetSpecializationByID(specializationID, companyID string) (specializations SpecializationDetail, err error)

	// GenerateLinkInvitation returns a link invitation
	GenerateLinkInvitation() (link string, err error)

	// UpdateSpecializationByID updates a specialization
	UpdateSpecializationByID(upsertUpdateSpec spec.UpsertUpdateSpecializationSpec) (id string, err error)
}

type specializationService struct {
	specializationRepo SpecializationRepository
	courseRepo         course.CourseRepository
	userRepo           users.UserRepository
	validate           *validator.Validate
}

func NewSpecializationService(
	specializationRepo SpecializationRepository,
	courseRepo course.CourseRepository,
	userRepo users.UserRepository,
) SpecializationService {
	return &specializationService{
		specializationRepo: specializationRepo,
		courseRepo:         courseRepo,
		userRepo:           userRepo,
		validate:           validator.New(),
	}
}

func (s *specializationService) Register(upsertSpecializationSpec spec.UpsertSpecializationSpec) (id string, err error) {
	err = s.validate.Struct(&upsertSpecializationSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	link := upsertSpecializationSpec.Invitation
	separateLink := strings.SplitAfter(link, "link=")[1]

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

	specialization, err = s.specializationRepo.FindInvitation(invitation)
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
			return specializations, exception.ErrInternalServer
		}

		countEmployee, err := s.specializationRepo.CountEmployee(companyID, spec.ID)
		if err != nil {
			return specializations, exception.ErrInternalServer
		}

		specializations = append(specializations, SpecializationDashboard{
			SpecializationID:   spec.ID,
			SpecializationName: spec.Name,
			AmountEmployee:     countEmployee,
			AmountCourse:       countCourse,
		})
	}

	return specializations, nil
}

func (s *specializationService) GenerateLinkInvitation() (link string, err error) {
	link = strings.Replace(uuid.New().String(), "-", "", -1)

	err = s.specializationRepo.CheckLinkInviation(link)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return link, nil
}

func (s *specializationService) GetSpecializationByID(specializationID, companyID string) (specializations SpecializationDetail, err error) {
	specialization, err := s.specializationRepo.FindSpecializationByID(specializationID, companyID)
	if err != nil {
		if err == exception.ErrSpecializationNotFound {
			return specializations, exception.ErrSpecializationNotFound
		}
		return specializations, exception.ErrInternalServer
	}

	// count course
	countCourse, err := s.specializationRepo.CountCourse(specialization.ID)
	if err != nil {
		return specializations, exception.ErrInternalServer
	}

	// count employee
	countEmployee, err := s.specializationRepo.CountEmployee(specialization.CompanyID, specialization.ID)
	if err != nil {
		return specializations, exception.ErrInternalServer
	}

	courses, err := s.courseRepo.FindAllCourseBySpecializationID(specialization.ID)
	if err != nil {
		if err == exception.ErrCourseNotFound {
			return specializations, exception.ErrCourseNotFound
		}

		return specializations, exception.ErrInternalServer
	}

	users, err := s.userRepo.FindAllUserBySpecializationID(specialization.ID)
	if err != nil {
		if err == exception.ErrEmployeeNotFound {
			return specializations, exception.ErrEmployeeNotFound
		}

		return specializations, exception.ErrInternalServer
	}

	specializations = SpecializationDetail{
		SpecializationID:   specialization.ID,
		CompanyID:          specialization.CompanyID,
		SpecializationName: specialization.Name,
		Invitation:         specialization.Invitation,
		AmountEmployee:     countEmployee,
		AmountCourse:       countCourse,
		Courses:            courses,
		Users:              users,
	}

	return specializations, nil
}

func (s *specializationService) AddCourseSpecialization(upsertCourseSpecializationSpec spec.UpsertCourseSpecializationSpec) (id string, err error) {
	err = s.validate.Struct(&upsertCourseSpecializationSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	id, err = s.specializationRepo.InsertCourseSpecialization(upsertCourseSpecializationSpec.CourseID, upsertCourseSpecializationSpec.SpecializationID)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}

func (s *specializationService) UpdateSpecializationByID(upsertUpdateSpec spec.UpsertUpdateSpecializationSpec) (id string, err error) {
	specialization, err := s.specializationRepo.FindSpecializationByID(upsertUpdateSpec.SpecializationID, upsertUpdateSpec.CompanyID)
	if err != nil {
		if err == exception.ErrSpecializationNotFound {
			return "", exception.ErrSpecializationNotFound
		}
		return "", exception.ErrInternalServer
	}

	newSpecilization := specialization.ModifySpecialization(upsertUpdateSpec.Name)

	id, err = s.specializationRepo.UpdateSpecialization(newSpecilization)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
