package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments/spec"

type RatingReviewsRequest struct {
	CourseID string
	UserID   string
	Rating   float32 `json:"rating"`
	Reviews  string  `json:"reviews"`
}

func (req *RatingReviewsRequest) ToSpecRatingReviews() *spec.UpsertRatingReviewSpec {
	return &spec.UpsertRatingReviewSpec{
		CourseID: req.CourseID,
		UserID:   req.UserID,
		Rating:   req.Rating,
		Reviews:  req.Reviews,
	}
}
