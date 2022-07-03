package response

type CreateNewSpecializationResponse struct {
	ID string `json:"id"`
}

func NewCreateNewSpecializationResponse(id string) CreateNewSpecializationResponse {
	return CreateNewSpecializationResponse{
		ID: id,
	}
}
