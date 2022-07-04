package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"

type CreateLoginResponse struct {
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
}

func NewCreateLoginResponse(auth auth.Auth) *CreateLoginResponse {
	var loginResponse CreateLoginResponse

	loginResponse.Token = auth.Token
	loginResponse.UserID = auth.UserID
	loginResponse.CompanyID = auth.CompanyID

	return &loginResponse
}
