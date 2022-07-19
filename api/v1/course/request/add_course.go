package request

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course/spec"
)

type CreateCourseRequest struct {
	ID          string
	CompanyID   string
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Modules     []modules `json:"modules"`
}

type modules struct {
	Title      string `json:"name"`
	YoutubeURL string `json:"youtube_url"`
	SlideURL   string `json:"slide_url"`
	Orders     int    `json:"orders"`
	Quizzes    []quiz `json:"quizzes"`
}

type quiz struct {
	Question       string   `json:"question"`
	Answer         string   `json:"answer"`
	MultipleChoice []string `json:"multiple_choice"`
}

func (req *CreateCourseRequest) ToSpec() *spec.UpsertCourseSpec {
	// Append modules and quizzes to spec
	var modules []spec.ModuleSpec
	for _, module := range req.Modules {
		var quizzes []spec.QuizSpec

		for _, quiz := range module.Quizzes {
			quizzes = append(quizzes, spec.QuizSpec{
				Question:       quiz.Question,
				Answer:         quiz.Answer,
				MultipleChoice: quiz.MultipleChoice,
			})
		}

		modules = append(modules, spec.ModuleSpec{
			Title:      module.Title,
			YoutubeURL: module.YoutubeURL,
			SlideURL:   module.SlideURL,
			Orders:     module.Orders,
			Quizzes:    quizzes,
		})
	}

	return &spec.UpsertCourseSpec{
		ID:          req.ID,
		CompanyID:   req.CompanyID,
		Title:       req.Title,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Modules:     modules,
	}
}
