package course

import (
	"fmt"
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/request"
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
	fmt.Println("masuk")
	CreateCourseRequest := new(request.CreateCourseRequest)
	if err := c.Bind(&CreateCourseRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *CreateCourseRequest.ToSpec()

	id, err := ctrl.service.Create(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}
