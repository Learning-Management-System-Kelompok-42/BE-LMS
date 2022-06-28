package response

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"

type GetAllUserDetailDashboardResp struct {
	ID          string                            `json:"id"`
	Name        string                            `json:"name"`
	Email       string                            `json:"email"`
	PhoneNumber string                            `json:"phone_number"`
	Address     string                            `json:"address"`
	Course      []users.CourseDetailDashboardUser `json:"course"`
	CreatedAt   string                            `json:"created_at"`
	UpdatedAt   string                            `json:"updated_at"`
}

func NewGetAllUserDetailDashboardResp(user users.UserDetailDashboard, courses []users.CourseDetailDashboardUser) GetAllUserDetailDashboardResp {
	var coursesList []users.CourseDetailDashboardUser
	for _, v := range courses {
		coursesList = append(coursesList, v)
	}

	return GetAllUserDetailDashboardResp{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Course:      coursesList,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
