package spec

type UpsertSpecializationSpec struct {
	CompanyID  string `validate:"required"`
	Name       string `validate:"required"`
	Invitation string `validate:"required"`
}
