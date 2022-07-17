package enrollments

import (
	"time"

	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type EnrollmentRepository interface {
	// InsertEnrollments insert a new enrollment
	InsertEnrollments(domain Domain) (id string, err error)

	// InsertRatingReviews insert a new rating and reviews
	InsertRatingReviews(domain Domain) (id string, err error)

	// FindAllModuleByCourseID find all modules by courseID
	FindAllModuleByCourseID(courseID string) (modules []module.Domain, err error)

	// FindEnrollmentByCourseIDUserID find enrollment by courseID and userID
	FindEnrollmentByCourseIDUserID(courseID string, userID string) (domain Domain, err error)

	// FindAllEnrollmentsByCourseID find all enrollment by course id
	FindAllEnrollmentsByCourseID(courseID string) (enrollments []RatingReviews, err error)

	// AVGRatingReviewsByCourseID count rating and reviews by course id
	AVGRatingReviewsByCourseID(courseID string) (avg float32, err error)

	// UpdateRatingReviews return id enrollments
	UpdateRatingReviews(domain Domain) (id string, err error)

	// CheckEnrollmentExist check if enrollment exist
	CheckEnrollmentExist(courseID, userID string) (err error)
}

type EnrollmentService interface {
	// CreateEnrollments Create insert a new enrollment
	CreateEnrollments(upsertEnrollSpec spec.UpsertEnrollSpec) (id string, err error)

	// CreateRatingReviews insert a new rating and reviews
	CreateRatingReviews(upsertRatingReviewSpec spec.UpsertRatingReviewSpec) (id string, err error)
}

type enrollmentService struct {
	enrollmentRepo EnrollmentRepository
	userModuleRepo userModules.UserModulesRepository
	validate       *validator.Validate
}

func NewEnrollmentService(enrollmentRepo EnrollmentRepository, userModuleRepo userModules.UserModulesRepository) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		userModuleRepo: userModuleRepo,
		validate:       validator.New(),
	}
}

func (s *enrollmentService) CreateEnrollments(upsertEnrollSpec spec.UpsertEnrollSpec) (id string, err error) {
	err = s.validate.Struct(&upsertEnrollSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	err = s.enrollmentRepo.CheckEnrollmentExist(upsertEnrollSpec.CourseID, upsertEnrollSpec.UserID)
	if err != nil {
		return "", exception.ErrEnrollmentAlreadyExist
	}

	newID := uuid.New().String()
	enrollAt := time.Now()
	status := false

	NewEnrollment := NewEnrollment(
		newID,
		upsertEnrollSpec.CourseID,
		upsertEnrollSpec.UserID,
		status,
		enrollAt,
	)

	id, err = s.enrollmentRepo.InsertEnrollments(NewEnrollment)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	modules, err := s.enrollmentRepo.FindAllModuleByCourseID(upsertEnrollSpec.CourseID)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	for _, v := range modules {
		newID := uuid.New().String()
		status := false
		var point int32 = 0

		newProgress := userModules.NewProggresCourse(
			newID,
			upsertEnrollSpec.UserID,
			upsertEnrollSpec.CourseID,
			v.ID,
			point,
			status,
		)

		_, err := s.userModuleRepo.InsertProgress(newProgress)
		if err != nil {
			return "", exception.ErrInternalServer
		}
	}

	return id, nil
}

func (s *enrollmentService) CreateRatingReviews(upsertRatingReviewSpec spec.UpsertRatingReviewSpec) (id string, err error) {
	err = s.validate.Struct(&upsertRatingReviewSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	oldEnrollment, err := s.enrollmentRepo.FindEnrollmentByCourseIDUserID(
		upsertRatingReviewSpec.CourseID,
		upsertRatingReviewSpec.UserID,
	)

	if err != nil {
		if err == exception.ErrEnrollmentNotFound {
			return "", exception.ErrEnrollmentNotFound
		}

		return "", exception.ErrInternalServer
	}

	newRatingReviews := oldEnrollment.NewRatingReviews(
		upsertRatingReviewSpec.Rating,
		upsertRatingReviewSpec.Reviews,
		true,
	)

	id, err = s.enrollmentRepo.UpdateRatingReviews(newRatingReviews)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
