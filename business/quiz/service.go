package quiz

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type QuizRepository interface {
	// Insert insert a new quiz
	Insert(quiz Domain) (id string, err error)
}

type QuizService interface {
	// Insert insert a new quiz
	Create(upsertQuizSpec spec.UpsertQuizSpec) (id string, err error)
}

type quizService struct {
	repo     QuizRepository
	validate *validator.Validate
}

func NewQuizService(repo QuizRepository) QuizService {
	return &quizService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *quizService) Create(upsertQuizSpec spec.UpsertQuizSpec) (id string, err error) {
	err = s.validate.Struct(&upsertQuizSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	newID := uuid.New().String()

	newQuiz := NewQuiz(
		newID,
		upsertQuizSpec.ModuleID,
		upsertQuizSpec.Question,
		upsertQuizSpec.Title,
		upsertQuizSpec.Answer,
		upsertQuizSpec.MultipleChoice,
	)

	id, err = s.repo.Insert(newQuiz)
	if err != nil {
		return "", exception.ErrInternalServer
	}

	return id, nil
}
