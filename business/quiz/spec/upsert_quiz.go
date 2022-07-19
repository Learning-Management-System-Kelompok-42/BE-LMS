package spec

type UpsertQuizSpec struct {
	ID             string
	ModuleID       string
	Title          string
	Question       string   `validate:"required"`
	Answer         string   `validate:"required"`
	MultipleChoice []string `validate:"required"`
}
