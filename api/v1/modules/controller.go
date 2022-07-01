package module

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/modules/response"
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
		} else if err == exception.ErrCourseNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}

func (ctrl *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	module, err := ctrl.service.GetByID(id)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.GetByIDModuleResponse(module)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) Update(c echo.Context) error {
	updateModuleRequest := new(request.UpdateModuleRequest)
	if err := c.Bind(updateModuleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *updateModuleRequest.ToSpecUpdate()

	id, err := ctrl.service.Update(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewCreateUpdateModuleResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(result))
}
