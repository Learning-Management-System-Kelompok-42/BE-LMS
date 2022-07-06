package course

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/course/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/course"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/s3"
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

func (ctrl *Controller) RegisterCourse(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	createCourseRequest := new(request.CreateCourseRequest)
	if err := c.Bind(&createCourseRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}
	createCourseRequest.CompanyID = companyID

	req := *createCourseRequest.ToSpec()

	id, err := ctrl.service.CreateCourse(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}

func (ctrl *Controller) GetDetailCourseDashboard(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	courseID := c.Param("courseID")

	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	course, err := ctrl.service.GetDetailCourseByIDDashboard(courseID)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetDetailCourseDashbordResp(course)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) GetAllCourseDashboard(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	course, err := ctrl.service.GetAllCourseDashboard(companyID)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetAllCourseDashboard(course)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	result, err := s3.UploadFileHelper(src, file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
