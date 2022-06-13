package spec

type UpsertCompanySpec struct {
	Name    string `validate:"required"`
	Address string `validate:"required"`
	Web     string `validate:"required"`
	Email   string `validate:"required,email"`
	Sector  string `validate:"required"`
	Logo    string
}
