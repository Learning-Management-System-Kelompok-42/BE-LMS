package specialization

import (
	"fmt"
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
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
	extract, _ := middleware.ExtractToken(c)

	createSpecializationRequest := new(request.CreateRequestSpecialization)
	createSpecializationRequest.CompanyID = extract.CompanyId

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

func (ctrl *Controller) GetInvitation(c echo.Context) error {
	invitation := c.QueryParam("link")

	spec, err := ctrl.service.GetInvitation(invitation)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetInvitationResponse(spec)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) GetAllSpecialization(c echo.Context) error {
	extract, _ := middleware.ExtractToken(c)

	spec, err := ctrl.service.GetAllSpecialization(extract.CompanyId)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	// result := response.NewGetAllSpecializationResponse(spec)

	return c.JSON(http.StatusOK, f.SuccessResponse(spec))
}

func (ctrl *Controller) GenerateLinkInvitation(c echo.Context) error {
	link, err := ctrl.service.GenerateLinkInvitation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := fmt.Sprintf("https://rubick.tech/invitation?link=%s", link)
	resp := response.NewGenerateLinkSpecializationResponse(result)

	return c.JSON(http.StatusOK, f.SuccessResponse(resp))
}
