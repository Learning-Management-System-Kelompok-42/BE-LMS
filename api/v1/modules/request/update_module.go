package request

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules/spec"
	quiz "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"
)

type UpdateModuleRequest struct {
	ID         string                `json:"id"`
	Title      string                `json:"title"`
	YoutubeURL string                `json:"youtube_url"`
	SlideURL   string                `json:"slide_url"`
	Orders     int                   `json:"orders"`
	Quizzes    []quiz.UpsertQuizSpec `json:"quizzes"`
}

func (req *UpdateModuleRequest) ToSpecUpdate() *spec.UpsertModuleSpec {
	return &spec.UpsertModuleSpec{
		ID:         req.ID,
		Title:      req.Title,
		YoutubeURL: req.YoutubeURL,
		SlideURL:   req.SlideURL,
		Orders:     req.Orders,
		Quizzes:    req.Quizzes,
	}
}
