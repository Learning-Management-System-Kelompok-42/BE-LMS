package course

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service course.CourseService
}

func NewController(service course.CourseService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createCourseRequest := new(request.CreateCourseRequest)
	if err := c.Bind(&createCourseRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createCourseRequest.ToSpec()

	id, err := ctrl.service.Create(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}

func (ctrl *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")
	// result, _ := middleware.ExtractToken(c)

	course, err := ctrl.service.GetByID(id)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetByIDCourseResponse(course)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) GetAllCourseDashboard(c echo.Context) error {
	extract, _ := middleware.ExtractToken(c)
	course, err := ctrl.service.GetAllCourseDashboard(extract.CompanyId)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetAllCourseDashboard(course)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}
