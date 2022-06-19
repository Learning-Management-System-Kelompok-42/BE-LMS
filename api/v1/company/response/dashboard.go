package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"

type GetDashboardResponse struct {
	UserID               string `json:"user_id"`
	CompanyID            string `json:"company_id"`
	NameAdmin            string `json:"name_admin"`
	NameCompany          string `json:"name_company"`
	AmountSpecialization int    `json:"amount_specialization"`
	AmountEmployee       int    `json:"amount_employee"`
}

func NewGetDashboardResponse(domain company.DashboardDomain) company.DashboardDomain {
	return company.DashboardDomain{
		UserID:               domain.UserID,
		CompanyID:            domain.CompanyID,
		NameAdmin:            domain.NameAdmin,
		NameCompany:          domain.NameCompany,
		AmountSpecialization: domain.AmountSpecialization,
		AmountEmployee:       domain.AmountEmployee,
	}
}
