package exception

import "errors"

var (
	// InternalServerError is the error code for internal server error
	ErrInternalServer = errors.New("internal server error")

	// NotFound
	ErrNotFound               = errors.New("not found")
	ErrEmployeeNotFound       = errors.New("employee not found")
	ErrCompanyNotFound        = errors.New("company not found")
	ErrAdminNotFound          = errors.New("admin not found")
	ErrCourseNotFound         = errors.New("course not found")
	ErrCommentNotFound        = errors.New("comment not found")
	ErrModuleNotFound         = errors.New("module not found")
	ErrVideoNotFound          = errors.New("video not found")
	ErrSlideNotFound          = errors.New("slide not found")
	ErrQuizNotFound           = errors.New("quiz not found")
	ErrSpecializationNotFound = errors.New("specialization not found")
	ErrEnrollmentNotFound     = errors.New("enrollment not found")
	ErrCourseAlreadyExist     = errors.New("course already exist")

	// NotAuthorized
	ErrNotAuthorized = errors.New("not authorized")

	// Other errors
	ErrWrongPassword        = errors.New("wrong password")
	ErrWrongEmail           = errors.New("wrong email")
	ErrWebExists            = errors.New("web already exists")
	ErrEmailExists          = errors.New("email already exists")
	ErrInvalidToken         = errors.New("invalid token")
	ErrDataNotFound         = errors.New("data not found")
	ErrInvalidRequest       = errors.New("invalid request body")
	ErrInvalidGenerateToken = errors.New("invalid generate token")
	ErrCantUploadImage      = errors.New("cant upload image")
)
