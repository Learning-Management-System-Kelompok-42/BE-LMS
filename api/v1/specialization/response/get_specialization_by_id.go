package response

import (
	course "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/response"
	users "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
)

type GetSpecializationByIDResponse struct {
	ID         string                         `json:"id"`
	CompanyID  string                         `json:"company_id"`
	Name       string                         `json:"name"`
	Invitation string                         `json:"invitation"`
	Courses    []course.GetAllCourseDashboard `json:"courses"`
	Users      []users.GetAllUsersResponse    `json:"users"`
}

func NewGetSpecializationByIDResponse(specialization specialization.SpecializationDetail) GetSpecializationByIDResponse {
	return GetSpecializationByIDResponse{
		ID:         specialization.SpecializationID,
		CompanyID:  specialization.CompanyID,
		Name:       specialization.SpecializationName,
		Invitation: specialization.Invitation,
		Courses:    course.NewGetAllCourseDashboard(specialization.Courses),
		Users:      users.NewGetAllUsersReponse(specialization.Users),
	}
}
