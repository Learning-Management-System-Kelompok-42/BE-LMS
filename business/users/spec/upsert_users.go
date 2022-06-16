package spec

type UpsertUsersSpec struct {
	CompanyID        string `validate:"required"`
	SpecializationID string `validate:"required"`
	Role             string `validate:"required"`
	FullName         string `validate:"required"`
	Email            string `validate:"required,email"`
	Password         string `validate:"required,min=5"`
	Phone            string `validate:"required"`
	Address          string `validate:"required"`
}
