package company

import (
	"fmt"
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/company/request"
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/company"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service domain.CompanyService
}

func NewController(service domain.CompanyService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createCompanyRequest := new(request.CreateRequestCompany)

	if err := c.Bind(&createCompanyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := *createCompanyRequest.ToSpecCreateCompany()

	fmt.Println("request = ", req)

	id, err := ctrl.service.Register(req)

	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, id)
}
