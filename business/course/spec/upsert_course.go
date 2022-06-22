package spec

type UpsertCourseSpec struct {
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	Thumbnail   string    `validate:"required"`
	Modules     []modules `validate:"required"`
}

// buat 1 1 struct
type modules struct {
	CourseID   string
	Title      string
	YoutubeURL string
	SlideURL   string
	Orders     int
	Quizzes    []quiz
}

type quiz struct {
	ModuleID       string
	Question       string
	Answer         string
	MultipleChoice []string
}
