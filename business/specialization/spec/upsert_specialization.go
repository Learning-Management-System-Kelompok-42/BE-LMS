package spec

type UpsertSpecializationSpec struct {
	CompanyID  string `validate:"required"`
	Name       string `validate:"required"`
	Invitation string `validate:"required"`
}

type UpsertCourseSpecializationSpec struct {
	SpecializationID string `validate:"required"`
	CourseID         string `validate:"required"`
}

type UpsertUpdateSpecializationSpec struct {
	SpecializationID string `validate:"required"`
	CompanyID        string `validate:"required"`
	Name             string `validate:"required"`
}
