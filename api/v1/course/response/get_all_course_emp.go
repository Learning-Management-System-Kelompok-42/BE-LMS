package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"

type GetAllCourseResp struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Progress    int64  `json:"progress"`
	Score       int64  `json:"score"`
}

func NewGetAllCourseResp(domain []course.ProgressCourse) []GetAllCourseResp {
	var courses []GetAllCourseResp
	for _, course := range domain {
		courses = append(courses, GetAllCourseResp{
			ID:          course.ID,
			Title:       course.Title,
			Thumbnail:   course.Thumbnail,
			Description: course.Description,
			Progress:    course.Proggress,
			Score:       course.Score,
		})
	}
	return courses
}
