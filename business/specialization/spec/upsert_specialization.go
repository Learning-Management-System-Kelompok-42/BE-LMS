package spec

type UpsertSpecializationSpec struct {
	Name       string `validate:"required"`
	Invitation string `validate:"required"`
}
