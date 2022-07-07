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
	AmountSpecialization string
	AmountEmployee       string
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
