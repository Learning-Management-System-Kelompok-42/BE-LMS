package users

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users/response"
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	r "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service domain.UserService
}

func NewController(service domain.UserService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	createUserRequest := new(request.CreateRequestUser)

	if err := c.Bind(&createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req := *createUserRequest.ToSpecCreateUsers()

	id, err := ctrl.service.Register(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, r.BadRequestResponse(err.Error()))
		} else if err == exception.ErrEmailExists {
			return c.JSON(http.StatusConflict, r.ConflictResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewCreateNewUserResponse(id)

	return c.JSON(http.StatusCreated, r.CreateSuccessResponse(result))
}

func (ctrl *Controller) GetUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := ctrl.service.GetUserByID(id)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
