package company

import "time"

type Domain struct {
	ID        string
	Name      string
	Address   string
	Web       string
	Sector    string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DashboardDomain struct {
	UserID               string
	CompanyID            string
	NameAdmin            string
	NameCompany          string
	AmountSpecialization int64
	AmountEmployee       int64
	AmountCourse         int64
	Specialization       []DashboardCompanySpecialization
	Course               []DashboardCompanyCourse
}

/*
	struct DashboardCompanySpecialization to be used in dashboard company page
	to show data specialization with highest amount of employee
**/
type DashboardCompanySpecialization struct {
	ID             string
	CompanyID      string
	Name           string
	Invitation     string
	AmountEmployee int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

/*
	struct DashboardCompanyCourse to be used in dashboard company page
	to show data course with highest amount of employee
**/
type DashboardCompanyCourse struct {
	ID             string
	CompanyID      string
	Title          string
	AmountEmployee int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

/*
	struct DashboardCompanyAdmin to be used in dashboard company page
	to show data admin name and company name
**/
type DashboardCompanyAdmin struct {
	UserID      string
	CompanyID   string
	NameAdmin   string
	NameCompany string
}

func NewCompany(id, name, address, web, Sector, Logo string) Domain {
	return Domain{
		ID:      id,
		Name:    name,
		Address: address,
		Web:     web,
		Sector:  Sector,
		Logo:    Logo,
	}
}

func (old *Domain) ModifyCompany(name, address, web, Sector, Logo string) Domain {
	return Domain{
		ID:        old.ID,
		Name:      name,
		Address:   address,
		Web:       web,
		Sector:    Sector,
		Logo:      Logo,
		CreatedAt: old.CreatedAt,
		UpdatedAt: old.UpdatedAt,
	}
}
