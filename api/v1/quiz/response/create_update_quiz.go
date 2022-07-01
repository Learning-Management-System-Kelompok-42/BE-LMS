package response

type CreateUpdateQuizResponse struct {
	ID string `json:"id"`
}

func NewCreateUpdateQuizResponse(id string) CreateUpdateQuizResponse {
	return CreateUpdateQuizResponse{
		ID: id,
	}
}
