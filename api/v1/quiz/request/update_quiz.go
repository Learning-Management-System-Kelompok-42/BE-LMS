package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"

type UpdateQuizRequest struct {
	ID             string   `json:"id"`
	ModuleID       string   `json:"module_id"`
	Title          string   `json:"title"`
	Question       string   `json:"question"`
	MultipleChoice []string `json:"multiple_choice"`
	Answer         string   `json:"answer"`
}

func (req *UpdateQuizRequest) ToSpecUpdate() *spec.UpsertQuizSpec {
	return &spec.UpsertQuizSpec{
		ID:             req.ID,
		ModuleID:       req.ModuleID,
		Title:          req.Title,
		Question:       req.Question,
		MultipleChoice: req.MultipleChoice,
		Answer:         req.Answer,
	}
}
