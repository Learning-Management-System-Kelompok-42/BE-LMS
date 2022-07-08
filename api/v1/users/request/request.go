package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"

type CreateRequestUser struct {
	CompanyID        string `json:"company_id"`
	SpecializationID string `json:"specialization_id"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number"`
	Address          string `json:"address"`
}

func (req *CreateRequestUser) ToSpecCreateUsers() *spec.UpsertUsersSpec {
	return &spec.UpsertUsersSpec{
		CompanyID:        req.CompanyID,
		SpecializationID: req.SpecializationID,
		FullName:         req.FullName,
		Email:            req.Email,
		Password:         req.Password,
		Phone:            req.PhoneNumber,
		Address:          req.Address,
	}
}
