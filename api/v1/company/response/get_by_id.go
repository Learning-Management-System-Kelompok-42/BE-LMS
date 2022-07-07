package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"

type GetCompanyProfileResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Sector    string `json:"sector"`
	Web       string `json:"website"`
	Logo      string `json:"logo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewGetCompanyProfileResponse(domain company.Domain) GetCompanyProfileResponse {
	return GetCompanyProfileResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Address:   domain.Address,
		Sector:    domain.Sector,
		Web:       domain.Web,
		Logo:      domain.Logo,
		CreatedAt: domain.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
