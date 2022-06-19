package specialization

type Domain struct {
	ID         string
	CompanyID  string
	Name       string
	Invitation string
}

func NewSpecialization(id, companyId, name, invitation string) Domain {
	return Domain{
		ID:         id,
		CompanyID:  companyId,
		Name:       name,
		Invitation: invitation,
	}
}
