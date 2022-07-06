package enrollments

import (
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
)

type Controller struct {
	service enrollments.EnrollmentService
}

func NewController(service enrollments.EnrollmentService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetAllEnrollmentsByCourseID(courseID string) (enrollments []enrollments.Domain, err error) {
	return enrollments, nil
}
