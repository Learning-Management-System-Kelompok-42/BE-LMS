package quiz

import (
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Quiz struct {
	ID        string `gorm:"primaryKey;size:200"`
	ModuleID  string `gorm:"size:200"`
	Title     string
	Question  string
	Options   pq.StringArray `gorm:"type:text[]"`
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (quizs *Quiz) ToDomain() quiz.Domain {
	return quiz.Domain{
		ID:             quizs.ID,
		ModuleID:       quizs.ModuleID,
		Title:          quizs.Title,
		Question:       quizs.Question,
		MultipleChoice: quizs.Options,
		Answer:         quizs.Answer,
		CreatedAt:      quizs.CreatedAt,
		UpdatedAt:      quizs.UpdatedAt,
	}
}

func (quizs *Quiz) ToDomainInBatch() course.QuizResp {
	return course.QuizResp{
		QuizID:         quizs.ID,
		ModuleID:       quizs.ModuleID,
		Question:       quizs.Question,
		MultipleChoice: quizs.Options,
		Answer:         quizs.Answer,
	}
}

func FromDomain(quiz quiz.Domain) Quiz {
	return Quiz{
		ID:        quiz.ID,
		ModuleID:  quiz.ModuleID,
		Title:     quiz.Title,
		Question:  quiz.Question,
		Options:   quiz.MultipleChoice,
		Answer:    quiz.Answer,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

func ToDomainInBatch(quizs []Quiz) []course.QuizResp {
	var domains []course.QuizResp
	for _, quiz := range quizs {
		domains = append(domains, quiz.ToDomainInBatch())
	}
	return domains
}
