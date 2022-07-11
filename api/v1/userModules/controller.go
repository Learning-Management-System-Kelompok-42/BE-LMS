package userModules

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/userModules/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/userModules/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/userModules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userModules.UserModulesService
}

func NewController(service userModules.UserModulesService) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) CreateProggress(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	courseID := c.Param("courseID")
	employeeID := c.Param("employeeID")

	if extract.UserId != employeeID {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	requestProggress := new(request.ProgressCourseRequest)
	if err := c.Bind(requestProggress); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	requestProggress.CourseID = courseID
	requestProggress.UserID = employeeID

	req := *requestProggress.ToSpec()

	id, err := ctrl.service.CreateProgress(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrProgressAlreadyExist {
			return c.JSON(http.StatusConflict, f.ConflictResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	resp := response.NewProgressCourseResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(resp))
}
