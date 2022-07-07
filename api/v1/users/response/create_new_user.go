package response

type CreateUpdateUserResponse struct {
	ID string `json:"id"`
}

func NewCreateNewUserResponse(id string) CreateUpdateUserResponse {
	return CreateUpdateUserResponse{
		ID: id,
	}
}

func NewCreateUpdateUserResponse(id string) CreateUpdateUserResponse {
	return CreateUpdateUserResponse{
		ID: id,
	}
}
