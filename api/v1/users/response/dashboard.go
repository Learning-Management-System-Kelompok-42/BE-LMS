package response

import (
	"sort"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
)

type dashboardUserResponse struct {
	UserID             string              `json:"user_id"`
	FullName           string              `json:"full_name"`
	SpecializationName string              `json:"specialization_name"`
	Course             int64               `json:"course"`
	CourseCompleted    int64               `json:"course_completed"`
	CourseInProgress   int64               `json:"course_in_progress"`
	TopCoursesProgress []topCourseProgress `json:"top_courses"`
	TopCourse7Days     []lastCourseOpen    `json:"top_course_7_days"`
}

type topCourseProgress struct {
	CourseID  string
	Thumbnail string
	Title     string
	Progress  int64
}

type lastCourseOpen struct {
	CourseID string
	Title    string
}

func DashboardUserResponse(domain users.DashboardEmployee) dashboardUserResponse {
	var coursesTop []topCourseProgress
	var coursesLastOpen []lastCourseOpen

	for _, course := range domain.TopCourseProgress {
		coursesTop = append(coursesTop, topCourseProgress{
			CourseID:  course.CourseID,
			Thumbnail: course.Thumbnail,
			Title:     course.Title,
			Progress:  course.Progress,
		})
	}

	if len(coursesTop) >= 2 {
		// sorting by progress
		sort.Slice(coursesTop, func(i, j int) bool {
			return coursesTop[i].Progress > coursesTop[j].Progress
		})

		// Get 1 course with highest progress
		coursesTop = coursesTop[:2]
	}

	for _, course := range domain.TopCourseOften7Days {
		coursesLastOpen = append(coursesLastOpen, lastCourseOpen{
			CourseID: course.CourseID,
			Title:    course.Title,
		})
	}

	return dashboardUserResponse{
		UserID:             domain.DetailEmployee.UserID,
		FullName:           domain.DetailEmployee.FullName,
		SpecializationName: domain.DetailEmployee.SpecializationName,
		Course:             domain.AmountCourse,
		CourseCompleted:    domain.AmountCourseCompleted,
		CourseInProgress:   domain.AmountCourseIncomplete,
		TopCoursesProgress: coursesTop,
		TopCourse7Days:     coursesLastOpen,
	}
}
