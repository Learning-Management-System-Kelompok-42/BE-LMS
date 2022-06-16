package auth

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/auth/response"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/auth"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service auth.AuthService
}

func NewController(service auth.AuthService) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Login(c echo.Context) error {
	requestLogin := new(request.CreateAuthRequest)
	if err := c.Bind(requestLogin); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *requestLogin.ToSpecAuth()

	user, err := ctrl.service.LoginUser(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		} else if err == exception.ErrWrongPassword {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		} else if err == exception.ErrInvalidGenerateToken {
			return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.CreateLoginResponse(*user)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}
