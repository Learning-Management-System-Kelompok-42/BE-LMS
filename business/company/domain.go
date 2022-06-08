package company

type Company struct {
	ID      string
	Name    string
	Address string
	Web     string
	Email   string
	Sector  string
	Logo    string
}

func NewCompany(id, name, address, web, email, Sector, Logo string) Company {
	return Company{
		ID:      id,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}

func (old *Company) ModifyCompany(name, address, web, email, Sector, Logo string) Company {
	return Company{
		ID:      old.ID,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}
