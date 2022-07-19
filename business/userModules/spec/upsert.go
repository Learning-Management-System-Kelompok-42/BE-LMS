package spec

type UpsertProgressSpec struct {
	UserID   string `validate:"required"`
	CourseID string `validate:"required"`
	ModuleID string `validate:"required"`
	Point    int32  `validate:"required"`
	Status   bool   `validate:"required"`
}
