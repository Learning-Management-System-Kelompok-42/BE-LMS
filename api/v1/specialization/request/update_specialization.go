package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"

type UpdateSpecializationRequest struct {
	SpecializationID string
	CompanyID        string
	Name             string `json:"name"`
}

func (req *UpdateSpecializationRequest) ToSpec() *spec.UpsertUpdateSpecializationSpec {
	return &spec.UpsertUpdateSpecializationSpec{
		SpecializationID: req.SpecializationID,
		CompanyID:        req.CompanyID,
		Name:             req.Name,
	}
}
