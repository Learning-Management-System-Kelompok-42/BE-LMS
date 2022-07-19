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

	// Update update a quiz
	Update(quiz Domain) (id string, err error)

	// GetByID get a quiz by id
	FindByID(id string) (quiz Domain, err error)
}

type QuizService interface {
	// Insert insert a new quiz
	Create(upsertQuizSpec spec.UpsertQuizSpec) (id string, err error)

	// Update update a quiz
	Update(upsertQuizSpec spec.UpsertQuizSpec) (id string, err error)

	// GetByID get a quiz by id
	GetByID(id string) (quiz Domain, err error)
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

func (s *quizService) Update(upsertQuizSpec spec.UpsertQuizSpec) (id string, err error) {
	err = s.validate.Struct(&upsertQuizSpec)
	if err != nil {
		return "", exception.ErrInvalidRequest
	}

	quiz, err := s.GetByID(upsertQuizSpec.ID)
	if err != nil {
		if err == exception.ErrNotFound {
			return id, exception.ErrNotFound
		}

		return id, exception.ErrInternalServer
	}

	updateQuiz := quiz.ModifyQuiz(
		upsertQuizSpec.Title,
		upsertQuizSpec.Question,
		upsertQuizSpec.Answer,
		upsertQuizSpec.MultipleChoice,
	)

	id, err = s.repo.Update(updateQuiz)
	if err != nil {
		if err == exception.ErrNotFound {
			return id, exception.ErrNotFound
		}

		return id, exception.ErrInternalServer
	}

	return id, nil
}

func (s *quizService) GetByID(id string) (quiz Domain, err error) {
	quiz, err = s.repo.FindByID(id)
	if err != nil {
		if err == exception.ErrNotFound {
			return quiz, exception.ErrNotFound
		}

		return quiz, exception.ErrInternalServer
	}

	return quiz, nil
}
