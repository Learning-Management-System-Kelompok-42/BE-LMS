package quiz

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) quiz.QuizRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(quiz quiz.Domain) (id string, err error) {
	newQuiz := FromDomain(quiz)
	err = repo.db.Create(&newQuiz).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newQuiz.ID

	return id, nil
}
