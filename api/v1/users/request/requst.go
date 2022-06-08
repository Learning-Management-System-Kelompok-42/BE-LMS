package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users/spec"

type CreateRequestUser struct {
	CompanyID   string `json:"company_id"`
	Role        string `json:"role"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	LevelAccess string `json:"level_access"`
}

func (req *CreateRequestUser) ToSpecCreateUsers() *spec.UpsertUsersSpec {
	return &spec.UpsertUsersSpec{
		CompanyID:   req.CompanyID,
		Role:        req.Role,
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Phone:       req.PhoneNumber,
		Address:     req.Address,
		LevelAccess: req.LevelAccess,
	}
}
