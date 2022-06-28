package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"

type GetByIDCourseResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewGetByIDCourseResponse(domain course.Domain) GetByIDCourseResponse {
	return GetByIDCourseResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Thumbnail:   domain.Thumbnail,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   domain.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
