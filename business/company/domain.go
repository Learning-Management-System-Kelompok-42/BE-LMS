package company

type Domain struct {
	ID      string
	Name    string
	Address string
	Web     string
	Email   string
	Sector  string
	Logo    string
}

func NewCompany(id, name, address, web, email, Sector, Logo string) Domain {
	return Domain{
		ID:      id,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}

func (old *Domain) ModifyCompany(name, address, web, email, Sector, Logo string) Domain {
	return Domain{
		ID:      old.ID,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}
