package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company/spec"

type CreateRequestCompany struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Web     string `json:"web"`
	Email   string `json:"email"`
	Sector  string `json:"sector"`
	Logo    string `json:"logo"`
}

func (req *CreateRequestCompany) ToSpecCreateCompany() *spec.UpsertCompanySpec {
	return &spec.UpsertCompanySpec{
		Name:    req.Name,
		Address: req.Address,
		Web:     req.Web,
		Email:   req.Email,
		Sector:  req.Sector,
		Logo:    req.Logo,
	}
}
