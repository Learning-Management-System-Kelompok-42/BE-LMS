package spec

type UpsertRequestFeat struct {
	UserID      string `validate:"required"`
	CompanyID   string `validate:"required"`
	RequestType string `validate:"required"`
	Title       string `validate:"required"`
	Reason      string `validate:"required"`
	RequestDate string `validate:"omitempty,required"`
	RequestTime string `validate:"omitempty,required"`
}
