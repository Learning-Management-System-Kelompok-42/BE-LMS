package enrollments

import "time"

type Domain struct {
	ID         string
	CourseID   string
	UserID     string
	Rating     float32
	Reviews    string
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

func NewEnrollment(id, courseID, userID, reviews string, rating float32, enrolledAt time.Time) Domain {
	return Domain{
		ID:         id,
		CourseID:   courseID,
		UserID:     userID,
		Rating:     float32(rating),
		Reviews:    reviews,
		EnrolledAt: enrolledAt,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}