package enrollments

import (
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

	return sum, nil
}

func (repo *postgreSQLRepository) InsertEnrollments(domain enrollments.Domain) (id string, err error) {
	newEnrollments := FromDomain(domain)

	err = repo.db.Create(&newEnrollments).Error
	if err != nil {
		return id, exception.ErrInternalServer
	}

	id = newEnrollments.ID

	return id, nil
}

func (repo *postgreSQLRepository) CheckEnrollmentExist(courseID string, userID string) (err error) {
	var enrollments Enrollments
	err = repo.db.Table("enrollments").Where("course_id = ? AND user_id = ?", courseID, userID).First(&enrollments).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return exception.ErrInternalServer
	}

	return exception.ErrEnrollmentAlreadyExist
}

func (repo *postgreSQLRepository) InsertRatingReviews(domain enrollments.Domain) (id string, err error) {
	newRatingReview := FromDomain(domain)

	err = repo.db.Save(&newRatingReview).Error
	if err != nil {
		return id, exception.ErrInternalServer
	}

	id = newRatingReview.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindEnrollmentByCourseIDUserID(courseID string, userID string) (domain enrollments.Domain, err error) {
	err = repo.db.Table("enrollments").Where("course_id = ? AND user_id = ?", courseID, userID).First(&domain).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain, exception.ErrEnrollmentNotFound
		}
		return domain, exception.ErrInternalServer
	}

	return domain, nil
}
