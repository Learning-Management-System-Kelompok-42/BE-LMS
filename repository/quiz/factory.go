package quiz

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/util"
)

func RepositoryFactory(dbConnection *util.DatabaseConnection) quiz.QuizRepository {
	var quizRepository quiz.QuizRepository

	if dbConnection.Driver == util.PostgreSQL {
		quizRepository = NewPostgreSQLRepository(dbConnection.PostgreSQL)
	}

	return quizRepository
}
