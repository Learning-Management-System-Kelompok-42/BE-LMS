package specialization

type Domain struct {
	ID         string
	Name       string
	Invitation string
}

func NewSpecialization(id, name, invitation string) Domain {
	return Domain{
		ID:         id,
		Name:       name,
		Invitation: invitation,
	}
}
