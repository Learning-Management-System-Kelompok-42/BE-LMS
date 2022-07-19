package spec

type UpsertAuthSpec struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
