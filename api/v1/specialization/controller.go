package specialization

import (
	"fmt"
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
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
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	createSpecializationRequest := new(request.CreateRequestSpecialization)
	createSpecializationRequest.CompanyID = companyID

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
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	spec, err := ctrl.service.GetAllSpecialization(companyID)
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
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	link, err := ctrl.service.GenerateLinkInvitation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := fmt.Sprintf("https://rubick.tech/invitation?link=%s", link)
	resp := response.NewGenerateLinkSpecializationResponse(result)

	return c.JSON(http.StatusOK, f.SuccessResponse(resp))
}

func (ctrl *Controller) GetDetailSpecialization(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	specializationID := c.Param("specializationID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	specialization, err := ctrl.service.GetSpecializationByID(specializationID, companyID)
	if err != nil {
		if err == exception.ErrSpecializationNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		} else if err == exception.ErrCourseNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		} else if err == exception.ErrEmployeeNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetSpecializationByIDResponse(specialization)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (ctrl *Controller) RegisterCourseSpecialization(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	SpecializationID := c.Param("specializationID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	createRequestRegister := new(request.CreateRequestCourseSpecialization)
	if err := c.Bind(&createRequestRegister); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	createRequestRegister.SpecializationID = SpecializationID

	req := *createRequestRegister.ToSpec()

	id, err := ctrl.service.AddCourseSpecialization(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrCourseAlreadyExist {
			return c.JSON(http.StatusConflict, f.ConflictResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	resp := response.NewCreateNewSpecializationResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(resp))
}

func (ctrl *Controller) UpdateSpecialization(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	SpecializationID := c.Param("specializationID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	updateRequest := new(request.UpdateSpecializationRequest)
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	updateRequest.SpecializationID = SpecializationID
	updateRequest.CompanyID = companyID

	req := *updateRequest.ToSpec()

	id, err := ctrl.service.UpdateSpecializationByID(req)
	if err != nil {
		if err == exception.ErrSpecializationNotFound {
			return c.JSON(http.StatusBadRequest, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	resp := response.NewCreateNewSpecializationResponse(id)

	return c.JSON(http.StatusOK, f.SuccessResponse(resp))
}
