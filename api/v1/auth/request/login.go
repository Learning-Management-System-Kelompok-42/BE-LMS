package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth/spec"

type CreateAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateAuthRequest) ToSpecAuth() *spec.UpsertAuthSpec {
	return &spec.UpsertAuthSpec{
		Email:    req.Email,
		Password: req.Password,
	}
}
