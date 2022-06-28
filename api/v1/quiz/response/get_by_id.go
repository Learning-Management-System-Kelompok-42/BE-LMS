package response

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
)

type GetByIDQuizResponse struct {
	ID             string    `json:"id"`
	ModuleID       string    `json:"module_id"`
	Title          string    `json:"title"`
	Question       string    `json:"question"`
	MultipleChoice []string  `json:"multiple_choice"`
	Answer         string    `json:"answer"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewGetByIDQuizResponse(quiz quiz.Domain) GetByIDQuizResponse {
	return GetByIDQuizResponse{
		ID:             quiz.ID,
		ModuleID:       quiz.ModuleID,
		Title:          quiz.Title,
		Question:       quiz.Question,
		MultipleChoice: quiz.MultipleChoice,
		Answer:         quiz.Answer,
		CreatedAt:      quiz.CreatedAt,
		UpdatedAt:      quiz.UpdatedAt,
	}
}
