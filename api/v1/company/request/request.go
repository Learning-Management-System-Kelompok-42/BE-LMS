package request

import (
	"mime/multipart"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"
)

type CreateRequestCompany struct {
	NameCompany    string         `form:"name_company"`
	AddressCompany string         `form:"address_company"`
	Sector         string         `form:"sector"`
	Website        string         `form:"website"`
	Logo           multipart.File `form:"logo"`
	NameAdmin      string         `form:"name_admin"`
	PhoneNumber    string         `form:"phone_number"`
	AddressAdmin   string         `form:"address_admin"`
	EmailAdmin     string         `form:"email_admin"`
	PasswordAdmin  string         `form:"password_admin"`
}

func (req *CreateRequestCompany) ToSpecCreateCompany() *spec.UpsertCompanySpec {
	return &spec.UpsertCompanySpec{
		NameCompany:    req.NameCompany,
		AddressCompany: req.AddressCompany,
		Sector:         req.Sector,
		Website:        req.Website,
		Logo:           req.Logo,
		NameAdmin:      req.NameAdmin,
		PhoneNumber:    req.PhoneNumber,
		AddressAdmin:   req.AddressAdmin,
		EmailAdmin:     req.EmailAdmin,
		PasswordAdmin:  req.PasswordAdmin,
	}
}
