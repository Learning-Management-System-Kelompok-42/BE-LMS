package response

type CreateNewUserResponse struct {
	ID string `json:"id"`
}

func NewCreateNewUserResponse(id string) CreateNewUserResponse {
	return CreateNewUserResponse{
		ID: id,
	}
}
