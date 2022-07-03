package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization/spec"

type CreateRequestCourseSpecialization struct {
	SpecializationID string
	CourseID         string `json:"course_id"`
}

func (req *CreateRequestCourseSpecialization) ToSpec() *spec.UpsertCourseSpecializationSpec {
	return &spec.UpsertCourseSpecializationSpec{
		SpecializationID: req.SpecializationID,
		CourseID:         req.CourseID,
	}
}
