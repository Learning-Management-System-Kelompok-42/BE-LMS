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

<<<<<<< Updated upstream
	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(id))
=======
	result := response.NewCreateUpdateQuizResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(result))
}

func (ctrl *Controller) Update(c echo.Context) error {
	createQuizRequest := new(request.UpdateQuizRequest)
	if err := c.Bind(createQuizRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createQuizRequest.ToSpecUpdate()

	id, err := ctrl.service.Update(req)
	if err != nil {
		if err == exception.ErrInvalidRequest {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}

		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewCreateUpdateQuizResponse(id)

	return c.JSON(http.StatusCreated, f.CreateSuccessResponse(result))
}

func (ctrl *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	quiz, err := ctrl.service.GetByID(id)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetByIDQuizResponse(quiz)

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
>>>>>>> Stashed changes
}
