package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments/spec"

type EnrollRequest struct {
	CourseID string
	UserID   string
}

func (req *EnrollRequest) ToSpec() *spec.UpsertEnrollSpec {
	return &spec.UpsertEnrollSpec{
		CourseID: req.CourseID,
		UserID:   req.UserID,
	}
}
