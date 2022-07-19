package response

type UpdateSpecializationNameResponse struct {
	UserID string `json:"user_id"`
}

func NewUpdateSpecializationNameResponse(id string) UpdateSpecializationNameResponse {
	return UpdateSpecializationNameResponse{
		UserID: id,
	}
}
