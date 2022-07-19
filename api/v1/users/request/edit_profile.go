package request

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"
)

type UpdateRequestUser struct {
	ID          string
	CompanyID   string
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	Avatar      string `form:"avatar"`
	FileName    string
}

func (req *UpdateRequestUser) ToSpecUpdateUsers() *spec.UpsertUpdateProfileSpec {
	return &spec.UpsertUpdateProfileSpec{
		ID:          req.ID,
		CompanyID:   req.CompanyID,
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Avatar:      req.Avatar,
		FileName:    req.FileName,
	}
}
