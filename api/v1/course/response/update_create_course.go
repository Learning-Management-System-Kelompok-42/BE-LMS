package response

type UpdateCreateCourseResponse struct {
	ID string `json:"id"`
}

func NewUpdateCreateCourseResponse(id string) UpdateCreateCourseResponse {
	return UpdateCreateCourseResponse{
		ID: id,
	}
}
