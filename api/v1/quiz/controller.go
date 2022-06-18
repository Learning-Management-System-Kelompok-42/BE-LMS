package quiz

import (
	"net/http"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/quiz/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/quiz"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service quiz.QuizService
}

func NewController(service quiz.QuizService) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Create(c echo.Context) error {
	createQuizRequest := new(request.CreateQuizRequest)
	if err := c.Bind(createQuizRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createQuizRequest.ToSpec()

	id, err := ctrl.service.Create(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
}
