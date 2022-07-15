package enrollments

import (
	"fmt"
<<<<<<< Updated upstream
=======
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"
>>>>>>> Stashed changes
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type EnrollmentRepository interface {
	// InsertEnrollments insert a new enrollment
	InsertEnrollments(domain Domain) (id string, err error)

<<<<<<< Updated upstream
=======
	// InsertRatingReviews insert a new rating and reviews
	InsertRatingReviews(domain Domain) (id string, err error)

	// FindAllModuleByCourseID find all modules by courseID
	FindAllModuleByCourseID(courseID string) (modules []module.Domain, err error)

	// FindEnrollmentByCourseIDUserID find enrollment by courseID and userID
	FindEnrollmentByCourseIDUserID(courseID string, userID string) (domain Domain, err error)

>>>>>>> Stashed changes
	// FindAllEnrollmentsByCourseID find all enrollment by course id
	FindAllEnrollmentsByCourseID(courseID string) (enrollments []RatingReviews, err error)

	// AVGRatingReviewsByCourseID count rating and reviews by course id
	AVGRatingReviewsByCourseID(courseID string) (avg float32, err error)

	// CheckEnrollmentExist check if enrollment exist
	CheckEnrollmentExist(courseID, userID string) (err error)
}

type EnrollmentService interface {
	// CreateEnrollments Create insert a new enrollment
	CreateEnrollments(upsertEnrollSpec spec.UpsertEnrollSpec) (id string, err error)

<<<<<<< Updated upstream
	// GetAllEnrollmentsByCourseID find all enrollment by course id
	// GetAllEnrollmentsByCourseID(courseID string) (enrollments []Domain, err error)
=======
	// CreateRatingReviews insert a new rating and reviews
	CreateRatingReviews(upsertRatingReviewSpec spec.UpsertRatingReviewSpec) (id string, err error)

>>>>>>> Stashed changes
}

type enrollmentService struct {
	enrollmentRepo EnrollmentRepository
<<<<<<< Updated upstream
	validator      *validator.Validate
=======
	userModuleRepo userModules.UserModulesRepository
	validate       *validator.Validate
>>>>>>> Stashed changes
}

func NewEnrollmentService(enrollmentRepo EnrollmentRepository, userModuleRepo userModules.UserModulesRepository) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
<<<<<<< Updated upstream
		validator:      validator.New(),
=======
		userModuleRepo: userModuleRepo,
		validate:       validator.New(),
>>>>>>> Stashed changes
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

		id, err := s.userModuleRepo.InsertProgress(newProgress)
		if err != nil {
			return "", exception.ErrInternalServer
		}

		fmt.Println("id = ", id)
	}

	return id, nil
}
