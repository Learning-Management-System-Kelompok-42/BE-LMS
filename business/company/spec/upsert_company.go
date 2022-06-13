package spec

import "mime/multipart"

type UpsertCompanySpec struct {
	NameCompany    string         `validate:"required"`
	AddressCompany string         `validate:"required"`
	Sector         string         `validate:"required"`
	Website        string         `validate:"required"`
	Logo           multipart.File `validate:"required"`
	NameAdmin      string         `validate:"required"`
	PhoneNumber    string         `validate:"required"`
	AddressAdmin   string         `validate:"required"`
	EmailAdmin     string         `validate:"required"`
	PasswordAdmin  string         `validate:"required"`
}
