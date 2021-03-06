package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"

type CreateQuizRequest struct {
	ModuleID       string   `json:"module_id"`
	Title          string   `json:"title"`
	Question       string   `json:"question"`
	Answer         string   `json:"answer"`
	MultipleChoice []string `json:"multiple_choice"`
}

func (req *CreateQuizRequest) ToSpec() *spec.UpsertQuizSpec {
	return &spec.UpsertQuizSpec{
		ModuleID:       req.ModuleID,
		Title:          req.Title,
		Question:       req.Question,
		Answer:         req.Answer,
		MultipleChoice: req.MultipleChoice,
	}
}
