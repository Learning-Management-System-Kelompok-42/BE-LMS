package specialization

type Domain struct {
	ID         string
	CompanyID  string
	Name       string
	Invitation string
}

type SpecializationDashboard struct {
	SpecializationID   string
	SpecializationName string
	AmountEmployee     int64
	AmountCourse       int64
}

func NewSpecialization(id, companyId, name, invitation string) Domain {
	return Domain{
		ID:         id,
		CompanyID:  companyId,
		Name:       name,
		Invitation: invitation,
	}
}
