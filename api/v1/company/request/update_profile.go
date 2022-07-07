package request

import (
	"mime/multipart"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"
)

type UpdateProfileCompanyRequest struct {
	CompanyID      string
	NameCompany    string         `form:"name_company"`
	AddressCompany string         `form:"address_company"`
	Sector         string         `form:"sector"`
	Website        string         `form:"website"`
	Logo           multipart.File `form:"logo"`
	FileName       string
}

func (req *UpdateProfileCompanyRequest) ToSpecUpdateCompany() *spec.UpsertProfileCompanySpec {
	return &spec.UpsertProfileCompanySpec{
		CompanyID:      req.CompanyID,
		NameCompany:    req.NameCompany,
		AddressCompany: req.AddressCompany,
		Sector:         req.Sector,
		Website:        req.Website,
		Logo:           req.Logo,
		FileName:       req.FileName,
	}
}
