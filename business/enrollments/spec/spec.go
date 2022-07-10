package spec

type UpsertEnrollSpec struct {
	CourseID string `validate:"required"`
	UserID   string `validate:"required"`
}

type UpsertRatingReviewSpec struct {
	CourseID string  `validate:"required"`
	UserID   string  `validate:"required"`
	Rating   float32 `validate:"required"`
	Reviews  string  `validate:"required"`
}
