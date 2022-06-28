package quiz

import (
	quizs "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) quizs.QuizRepository {
	return &postgreSQLRepository{db: db}
}

func (repo *postgreSQLRepository) Insert(quiz quizs.Domain) (id string, err error) {
	newQuiz := FromDomain(quiz)
	err = repo.db.Create(&newQuiz).Error

	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = newQuiz.ID

	return id, nil
}

func (repo *postgreSQLRepository) Update(quiz quizs.Domain) (id string, err error) {
	updateQuiz := FromDomain(quiz)

	err = repo.db.Where("id = ?", quiz.ID).Save(&updateQuiz).Error
	if err != nil {
		return "", exception.ErrInternalServer
	}

	id = updateQuiz.ID

	return id, nil
}

func (repo *postgreSQLRepository) FindByID(id string) (quiz quizs.Domain, err error) {
	returnQuiz := FromDomain(quizs.Domain{ID: id})

	err = repo.db.Where("id = ?", id).First(&returnQuiz).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return quiz, exception.ErrNotFound
		}

		return quiz, exception.ErrInternalServer
	}

	quiz = returnQuiz.ToDomain()

	return quiz, nil
}
