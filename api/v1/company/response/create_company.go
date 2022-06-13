package response

type CreateNewCompanyResponse struct {
	ID string `json:"id"`
}

func NewCreateNewCompanyResponse(id string) CreateNewCompanyResponse {
	return CreateNewCompanyResponse{
		ID: id,
	}
}
