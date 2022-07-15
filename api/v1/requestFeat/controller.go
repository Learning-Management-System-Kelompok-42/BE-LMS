package requestFeat

import (
	"fmt"
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/requestFeat/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service requestFeat.RequestFeatService
}

func NewController(service requestFeat.RequestFeatService) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) CreateRequestFeat(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	userID := c.Param("employeeID")

	if userID != extract.UserId {
		fmt.Println("userid = ", userID)
		fmt.Println("user id = ", extract.UserId)
		return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	createRequestFeatRequest := new(request.RequestFeatReq)
	if err := c.Bind(&createRequestFeatRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	createRequestFeatRequest.UserID = userID

	req := *createRequestFeatRequest.ToSpec()

	id, err := ctrl.service.CreateRequestFeat(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}
