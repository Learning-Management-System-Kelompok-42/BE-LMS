package spec

type UpsertCourseSpec struct {
	CompanyID   string       `validate:"required"`
	Title       string       `validate:"required"`
	Description string       `validate:"required"`
	Thumbnail   string       `validate:"required"`
	Modules     []ModuleSpec `validate:"required"`
}

// buat 1 1 struct
type ModuleSpec struct {
	CourseID   string
	Title      string
	YoutubeURL string
	SlideURL   string
	Orders     int
	Quizzes    []QuizSpec
}

type QuizSpec struct {
	ModuleID       string
	Question       string
	Answer         string
	MultipleChoice []string
}
