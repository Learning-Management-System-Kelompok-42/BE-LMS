package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"

type UpdateSpecializationNameRequest struct {
	SpecializationID string `json:"specialization_id"`
	CompanyID        string
	UserID           string
}

func (req *UpdateSpecializationNameRequest) ToSpec() *spec.UpsertUpdateSpecName {
	return &spec.UpsertUpdateSpecName{
		SpecializationID: req.SpecializationID,
		CompanyID:        req.CompanyID,
		UserID:           req.UserID,
	}
}
