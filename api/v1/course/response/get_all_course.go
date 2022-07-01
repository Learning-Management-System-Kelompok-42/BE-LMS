package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"

type GetAllCourseDashboard struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func NewGetAllCourseDashboard(domain []course.Domain) []GetAllCourseDashboard {
	var courses []GetAllCourseDashboard
	for _, course := range domain {
		courses = append(courses, GetAllCourseDashboard{
			ID:          course.ID,
			Title:       course.Title,
			Thumbnail:   course.Thumbnail,
			Description: course.Description,
			Rating:      4.5,
			CreatedAt:   course.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   course.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return courses
}
