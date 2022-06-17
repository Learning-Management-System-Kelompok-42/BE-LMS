package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"

type GetUserByID struct {
	ID               string `json:"id"`
	CompanyID        string `json:"company_id"`
	Role             string `json:"role"`
	SpecializationID string `json:"specialization_id"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number"`
	Address          string `json:"address"`
	LevelAccess      string `json:"level_access"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func FromDomainUser(user *users.Domain) GetUserByID {
	return GetUserByID{
		ID:               user.ID,
		CompanyID:        user.CompanyID,
		Role:             user.Role,
		SpecializationID: user.SpecializationID,
		FullName:         user.FullName,
		Email:            user.Email,
		Password:         user.Password,
		PhoneNumber:      user.PhoneNumber,
		Address:          user.Address,
		LevelAccess:      user.LevelAccess,
		CreatedAt:        user.CreatedAt.String(),
		UpdatedAt:        user.UpdatedAt.String(),
	}
}
