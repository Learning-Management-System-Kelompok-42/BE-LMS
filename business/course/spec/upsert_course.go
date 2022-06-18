package spec

type UpsertCourseSpec struct {
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	Thumbnail   string    `validate:"required"`
	Modules     []Modules `validate:"required"`
}

type Modules struct {
	CourseID   string
	Title      string `json:"name"`
	YoutubeURL string `json:"youtube_url"`
	SlideURL   string `json:"slide_url"`
	Orders     int    `json:"orders"`
	Quizzes    []Quiz `json:"quizzes"`
}

type Quiz struct {
	Question       string   `json:"question"`
	Answer         string   `json:"answer"`
	MultipleChoice []string `json:"multiple_choice"`
}
