package spec

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz/spec"

type UpsertModuleSpec struct {
	CourseID   string
	Title      string `validate:"required"`
	YoutubeURL string `validate:"required"`
	SlideURL   string `validate:"required"`
	Orders     int    `validate:"required"`
	Quizzes    []spec.UpsertQuizSpec
}
