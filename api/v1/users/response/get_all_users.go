package response

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
)

type GetAllUsersResponse struct {
	ID               string `json:"id"`
	CompanyID        string `json:"company_id"`
	Role             string `json:"specialization_name"`
	SpecializationID string `json:"specialization_id"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	Address          string `json:"address"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func FromDomainList(users users.Domain) GetAllUsersResponse {
	return GetAllUsersResponse{
		ID:               users.ID,
		CompanyID:        users.CompanyID,
		Role:             users.Role,
		SpecializationID: users.SpecializationID,
		FullName:         users.FullName,
		Email:            users.Email,
		PhoneNumber:      users.PhoneNumber,
		Address:          users.Address,
		CreatedAt:        users.CreatedAt.String(),
		UpdatedAt:        users.UpdatedAt.String(),
	}
}

func NewGetAllUsersReponse(users []users.Domain) []GetAllUsersResponse {
	list := []GetAllUsersResponse{}

	for _, user := range users {
		list = append(list, FromDomainList(user))
	}

	return list
}
