package response

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
)

type GetAllUserDetailDashboardResp struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Role        string             `json:"specialization_name"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phone_number"`
	Address     string             `json:"address"`
	Course      []getAllCourseResp `json:"course"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}

type getAllCourseResp struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func NewGetAllUserDetailDashboardResp(user users.UserDetailDashboard, courses []users.CourseDetailDashboardUser) GetAllUserDetailDashboardResp {
	var coursesList []getAllCourseResp
	for _, v := range courses {
		coursesList = append(coursesList, getAllCourseResp{
			ID:          v.ID,
			Name:        v.Name,
			Thumbnail:   v.Thumbnail,
			Description: v.Description,
			Rating:      v.Rating,
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return GetAllUserDetailDashboardResp{
		ID:          user.ID,
		Name:        user.Name,
		Role:        user.Role,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Course:      coursesList,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
