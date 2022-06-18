package quiz

import "time"

type Domain struct {
	ID             string
	ModuleID       string
	Title          string
	Question       string
	MultipleChoice []string
	Answer         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewQuiz(id, moduleID, question, title, answer string, multipleChoice []string) Domain {
	return Domain{
		ID:             id,
		ModuleID:       moduleID,
		Title:          title,
		Question:       question,
		MultipleChoice: multipleChoice,
		Answer:         answer,
	}
}

func (old *Domain) ModifyQuiz(title, question, answer string, multipleChoice []string) Domain {
	return Domain{
		ID:             old.ID,
		ModuleID:       old.ModuleID,
		Title:          title,
		Question:       question,
		MultipleChoice: multipleChoice,
		Answer:         answer,
	}
}
