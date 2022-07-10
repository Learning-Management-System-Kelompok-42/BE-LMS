package enrollments

import (
	"fmt"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type EnrollmentRepository interface {
	// Insert insert a new enrollment
	InsertEnrollments(domain Domain) (id string, err error)

	// FindAllEnrollmentsByCourseID find all enrollment by course id
	FindAllEnrollmentsByCourseID(courseID string) (enrollments []RatingReviews, err error)

	// CountRatingReviewsByCourseID count rating and reviews by course id
	AVGRatingReviewsByCourseID(courseID string) (avg float32, err error)

	// CheckEnrollmentExist check if enrollment exist
	CheckEnrollmentExist(courseID, userID string) (err error)
}

type EnrollmentService interface {
	// Create insert a new enrollment
	CreateEnrollments(upsertEnrollSpec spec.UpsertEnrollSpec) (id string, err error)

	// GetAllEnrollmentsByCourseID find all enrollment by course id
	// GetAllEnrollmentsByCourseID(courseID string) (enrollments []Domain, err error)
}

type enrollmentService struct {
	enrollmentRepo EnrollmentRepository
	validator      *validator.Validate
}

func NewEnrollmentService(enrollmentRepo EnrollmentRepository) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		validator:      validator.New(),
	}
}

func (s *enrollmentService) CreateEnrollments(upsertEnrollSpec spec.UpsertEnrollSpec) (id string, err error) {
	err = s.validator.Struct(&upsertEnrollSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	err = s.enrollmentRepo.CheckEnrollmentExist(upsertEnrollSpec.CourseID, upsertEnrollSpec.UserID)
	if err != nil {
		return "", exception.ErrEnrollmentAlreadyExist
	}

	newID := uuid.New().String()
	enrollAt := time.Now()
	fmt.Println("enrollAt: ", enrollAt)

	NewEnrollment := NewEnrollment(
		newID,
		upsertEnrollSpec.CourseID,
		upsertEnrollSpec.UserID,
		enrollAt,
	)

	id, err = s.enrollmentRepo.InsertEnrollments(NewEnrollment)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
