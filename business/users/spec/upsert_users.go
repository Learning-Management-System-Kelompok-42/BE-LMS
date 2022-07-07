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

type UpsertUpdateProfileSpec struct {
	ID          string `validate:"required"`
	CompanyID   string `validate:"required"`
	FullName    string `validate:"required"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required"`
	Address     string `validate:"required"`
	Avatar      string
	FileName    string
}

type UpsertUpdateSpecName struct {
	SpecializationID string `validate:"required"`
	CompanyID        string `validate:"required"`
	UserID           string `validate:"required"`
}
