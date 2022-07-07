package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"

type UpdatePasswordRequestUser struct {
	UserID      string
	CompanyID   string
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (req *UpdatePasswordRequestUser) ToSpecUpdatePassowrd() *spec.UpsertUpdatePassword {
	return &spec.UpsertUpdatePassword{
		UserID:      req.UserID,
		CompanyID:   req.CompanyID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}
