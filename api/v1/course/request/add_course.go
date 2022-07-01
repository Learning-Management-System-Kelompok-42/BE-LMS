package request

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course/spec"
)

type CreateCourseRequest struct {
<<<<<<< Updated upstream
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Thumbnail   string         `json:"thumbnail"`
	Modules     []spec.Modules `json:"modules"`
=======
	CompanyID   string
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Modules     []modules `json:"modules"`
>>>>>>> Stashed changes
}

// type modules struct {
// 	Title      string `json:"name"`
// 	YoutubeURL string `json:"youtube_url"`
// 	SlideURL   string `json:"slide_url"`
// 	Orders     int    `json:"orders"`
// 	Quizzes    []quiz `json:"quizzes"`
// }

// type quiz struct {
// 	Question       string   `json:"question"`
// 	Answer         string   `json:"answer"`
// 	MultipleChoice []string `json:"multiple_choice"`
// }

func (req *CreateCourseRequest) ToSpec() *spec.UpsertCourseSpec {
	return &spec.UpsertCourseSpec{
		CompanyID:   req.CompanyID,
		Title:       req.Title,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Modules:     req.Modules,
	}
}
