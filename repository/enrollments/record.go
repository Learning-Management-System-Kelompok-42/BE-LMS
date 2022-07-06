package enrollments

import (
	"time"

	enrollment "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	"gorm.io/gorm"
)

type Enrollments struct {
	ID         string  `gorm:"primaryKey;size:200;autoIncrement"`
	CourseID   string  `gorm:"size:200"`
	UserID     string  `gorm:"size:200"`
	Rating     float32 `gorm:"type:numeric(2,2)"`
	Reviews    string
	EnrolledAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type RatingReviews struct {
	ID      string
	Name    string
	Email   string
	Rating  float32
	Reviews string
}

func (enrollments *Enrollments) ToDomain() enrollment.Domain {
	return enrollment.Domain{
		ID:         enrollments.ID,
		CourseID:   enrollments.CourseID,
		UserID:     enrollments.UserID,
		Rating:     enrollments.Rating,
		Reviews:    enrollments.Reviews,
		EnrolledAt: enrollments.EnrolledAt,
		CreatedAt:  enrollments.CreatedAt,
		UpdatedAt:  enrollments.UpdatedAt,
	}
}

func (ratingReviews *RatingReviews) ToDomain() enrollment.RatingReviews {
	return enrollment.RatingReviews{
		ID:      ratingReviews.ID,
		Name:    ratingReviews.Name,
		Email:   ratingReviews.Email,
		Rating:  ratingReviews.Rating,
		Reviews: ratingReviews.Reviews,
	}
}

func FromDomain(enrollment enrollment.Domain) Enrollments {
	return Enrollments{
		ID:         enrollment.ID,
		CourseID:   enrollment.CourseID,
		UserID:     enrollment.UserID,
		Rating:     enrollment.Rating,
		Reviews:    enrollment.Reviews,
		EnrolledAt: enrollment.EnrolledAt,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  gorm.DeletedAt{},
	}
}

func ToDomainBatchList(enrollments []Enrollments) []enrollment.Domain {
	var enromlment []enrollment.Domain

	for _, enrollment := range enrollments {
		enromlment = append(enromlment, enrollment.ToDomain())
	}

	return enromlment
}

func ToDomainRatingList(ratingReviews []RatingReviews) []enrollment.RatingReviews {
	var enromlment []enrollment.RatingReviews

	for _, enrollment := range ratingReviews {
		enromlment = append(enromlment, enrollment.ToDomain())
	}

	return enromlment
}
