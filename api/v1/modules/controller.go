package module

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules/request"
	module "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/modules"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service module.ModuleService
}

func NewController(service module.ModuleService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createModuleRequest := new(request.CreateModuleRequest)
	if err := c.Bind(createModuleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createModuleRequest.ToSpec()

	id, err := ctrl.service.Create(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}
