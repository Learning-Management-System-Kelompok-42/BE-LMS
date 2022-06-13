package company

<<<<<<< Updated upstream
type Company struct {
=======
type Domain struct {
>>>>>>> Stashed changes
	ID      string
	Name    string
	Address string
	Web     string
	Email   string
	Sector  string
	Logo    string
}

<<<<<<< Updated upstream
func NewCompany(id, name, address, web, email, Sector, Logo string) Company {
	return Company{
=======
func NewCompany(id, name, address, web, email, Sector, Logo string) Domain {
	return Domain{
>>>>>>> Stashed changes
		ID:      id,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}

<<<<<<< Updated upstream
func (old *Company) ModifyCompany(name, address, web, email, Sector, Logo string) Company {
	return Company{
=======
func (old *Domain) ModifyCompany(name, address, web, email, Sector, Logo string) Domain {
	return Domain{
>>>>>>> Stashed changes
		ID:      old.ID,
		Name:    name,
		Address: address,
		Web:     web,
		Email:   email,
		Sector:  Sector,
		Logo:    Logo,
	}
}
