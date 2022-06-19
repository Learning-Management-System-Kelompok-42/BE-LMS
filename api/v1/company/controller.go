package company

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	r "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service company.CompanyService
}

func NewController(service company.CompanyService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createCompanyRequest := new(request.CreateRequestCompany)

	file, err := c.FormFile("logo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.BadRequestResponse(err.Error()))
	}

	formFile, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.BadRequestResponse(err.Error()))
	}

	if err := c.Bind(createCompanyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, r.BadRequestResponse(err.Error()))
	}

	req := *createCompanyRequest.ToSpecCreateCompany()

	req.Logo = formFile

	id, err := ctrl.service.Register(req)

	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, r.BadRequestResponse(err.Error()))
		} else if err == exception.ErrEmailExists {
			return c.JSON(http.StatusConflict, r.ConflictResponse(err.Error()))
		} else if err == exception.ErrWebExists {
			return c.JSON(http.StatusConflict, r.ConflictResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewCreateNewCompanyResponse(id)

	return c.JSON(http.StatusCreated, r.CreateSuccessResponse(result))
}

func (ctrl *Controller) GetDashboard(c echo.Context) error {
	companyID, _, _, _ := m.ExtractToken(c)

	company, err := ctrl.service.Dashboard(companyID)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetDashboardResponse(company)

	return c.JSON(http.StatusOK, r.SuccessResponse(result))
}
