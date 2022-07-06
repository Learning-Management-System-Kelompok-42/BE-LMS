package enrollments

import "github.com/go-playground/validator/v10"

type EnrollmentRepository interface {
	// Insert insert a new enrollment
	// InsertEnrollments(enrollment Domain) (id string, err error)

	// FindAllEnrollmentsByCourseID find all enrollment by course id
	FindAllEnrollmentsByCourseID(courseID string) (enrollments []RatingReviews, err error)
}

type EnrollmentService interface {
	// Create insert a new enrollment
	// CreateEnrollments(enrollment Domain) (id string, err error)

	// GetAllEnrollmentsByCourseID find all enrollment by course id
	// GetAllEnrollmentsByCourseID(courseID string) (enrollments []Domain, err error)
}

type enrollmentService struct {
	enrollmentRepo EnrollmentRepository
	validator      *validator.Validate
}

func NewEnrollmentService(enrollmentRepo EnrollmentRepository, validator *validator.Validate) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		validator:      validator,
	}
}
