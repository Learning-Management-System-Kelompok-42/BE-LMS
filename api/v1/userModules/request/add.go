package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules/spec"

type ProgressCourseRequest struct {
	UserID   string
	CourseID string
	ModuleID string `json:"module_id"`
	Point    int32  `json:"point"`
	Status   bool   `json:"status"`
}

func (req *ProgressCourseRequest) ToSpec() *spec.UpsertProgressSpec {
	return &spec.UpsertProgressSpec{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		ModuleID: req.ModuleID,
		Point:    req.Point,
		Status:   req.Status,
	}
}
