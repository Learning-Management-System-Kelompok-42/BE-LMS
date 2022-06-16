package specialization

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/specialization/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/specialization"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service specialization.SpecializationService
}

func NewController(service specialization.SpecializationService) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createSpecializationRequest := new(request.CreateRequestSpecialization)

	if err := c.Bind(&createSpecializationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createSpecializationRequest.ToSpec()

	id, err := ctrl.service.Register(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewCreateNewSpecializationResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(result))
}
