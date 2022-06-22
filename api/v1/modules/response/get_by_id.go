package response

import (
	"time"

	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
)

type GetByIDModuleResponse struct {
	ID         string    `json:"id"`
	CourseID   string    `json:"course_id"`
	YoutubeURL string    `json:"youtube_url"`
	SlideURL   string    `json:"slide_url"`
	Title      string    `json:"title"`
	Orders     int       `json:"orders"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewGetByIDModuleResponse(module module.Domain) GetByIDModuleResponse {
	return GetByIDModuleResponse{
		ID:         module.ID,
		CourseID:   module.CourseID,
		YoutubeURL: module.YoutubeURL,
		SlideURL:   module.SlideURL,
		Title:      module.Title,
		Orders:     module.Orders,
		CreatedAt:  module.CreatedAt,
		UpdatedAt:  module.UpdatedAt,
	}
}
