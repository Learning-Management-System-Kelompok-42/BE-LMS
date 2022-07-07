package enrollments

import (
	"fmt"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) enrollments.EnrollmentRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) FindAllEnrollmentsByCourseID(courseID string) (enrollments []enrollments.RatingReviews, err error) {
	var newEnrollments []RatingReviews
	result := repo.db.Table("enrollments").
		Select("enrollments.id, enrollments.rating, enrollments.reviews, users.full_name as name, users.email").
		Joins("JOIN users ON enrollments.user_id = users.id").
		Where("enrollments.course_id = ?", courseID).
		Find(&newEnrollments)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return enrollments, exception.ErrEnrollmentNotFound
		}
		return enrollments, exception.ErrInternalServer
	}

	enrollments = ToDomainRatingList(newEnrollments)

	return enrollments, nil
}

func (repo *postgreSQLRepository) AVGRatingReviewsByCourseID(courseID string) (avg float32, err error) {
	var newEnrollments []RatingReviews
	result := repo.db.Table("enrollments").
		Select("enrollments.rating").
		Joins("JOIN users ON enrollments.user_id = users.id").
		Where("enrollments.course_id = ?", courseID).
		Find(&newEnrollments)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return avg, exception.ErrEnrollmentNotFound
		}
		return avg, exception.ErrInternalServer
	}

	// Calculate average rating
	var sum float32
	for _, enrollment := range newEnrollments {
		sum += enrollment.Rating
	}

	fmt.Println("sum: ", sum)

	return sum, nil
}
