package enrollments

import "time"

type Domain struct {
	ID         string
	CourseID   string
	UserID     string
	Rating     float32
	Reviews    string
	Status     bool
	EnrolledAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type RatingReviews struct {
	ID      string
	Name    string
	Email   string
	Rating  float32
	Reviews string
}

func NewEnrollment(id, courseID, userID string, status bool, enrolledAt time.Time) Domain {
	return Domain{
		ID:         id,
		CourseID:   courseID,
		UserID:     userID,
		Status: status,
		EnrolledAt: enrolledAt,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

// func NewRatingReviews(rating float32, reviews string) Domain {
// 	return Domain{
// 		Rating:  rating,
// 		Reviews: reviews,
// 	}
// }

func (old *Domain) NewRatingReviews(rating float32, reviews string) Domain {
	return Domain{
		ID:         old.ID,
		UserID:     old.UserID,
		CourseID:   old.CourseID,
		Rating:     rating,
		Reviews:    reviews,
		EnrolledAt: old.EnrolledAt,
		CreatedAt:  old.CreatedAt,
		UpdatedAt:  old.UpdatedAt,
	}
}
