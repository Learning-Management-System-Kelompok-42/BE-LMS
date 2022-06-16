package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"

type CreateRequestSpecialization struct {
	Name       string `json:"name"`
	Invitation string `json:"invitation"`
}

func (req *CreateRequestSpecialization) ToSpec() *spec.UpsertSpecializationSpec {
	return &spec.UpsertSpecializationSpec{
		Name:       req.Name,
		Invitation: req.Invitation,
	}
}
