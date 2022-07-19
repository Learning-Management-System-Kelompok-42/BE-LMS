package enrollments

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/enrollments/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/enrollments"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
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

func (ctrl *Controller) CreateEnrollments(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	employeeID := c.Param("employeeID")
	courseID := c.Param("courseID")

	if employeeID != extract.UserId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	requestEnroll := new(request.EnrollRequest)

	requestEnroll.CourseID = courseID
	requestEnroll.UserID = employeeID

	req := *requestEnroll.ToSpec()

	id, err := ctrl.service.CreateEnrollments(req)
	if err != nil {
		if err == exception.ErrEnrollmentAlreadyExist {
			return c.JSON(http.StatusConflict, f.ConflictResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}

func (ctrl *Controller) CreateRatingReviews(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	employeeID := c.Param("employeeID")
	courseID := c.Param("courseID")

	if employeeID != extract.UserId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	createRatingReviews := new(request.RatingReviewsRequest)
	if err := c.Bind(createRatingReviews); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	createRatingReviews.CourseID = courseID
	createRatingReviews.UserID = employeeID

	req := *createRatingReviews.ToSpecRatingReviews()

	id, err := ctrl.service.CreateRatingReviews(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrEnrollmentNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}
