package response

type ProgressCourseResponse struct {
	ID string `json:"id"`
}

func NewProgressCourseResponse(id string) ProgressCourseResponse {
	return ProgressCourseResponse{
		ID: id,
	}
}
