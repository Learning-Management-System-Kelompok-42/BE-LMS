package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"

type GetInvitationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Invitation string `json:"invitation"`
}

func NewGetInvitationResponse(spec specialization.Domain) *GetInvitationResponse {
	var specializationResponse GetInvitationResponse

	specializationResponse.ID = spec.ID
	specializationResponse.Name = spec.Name
	specializationResponse.Invitation = spec.Invitation

	return &specializationResponse
}
