package users

import (
	"net/http"

	m "github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/middleware"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users/request"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/api/v1/users/response"
	domain "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/users"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/exception"
	r "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"
	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/s3"
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

func (ctrl *Controller) GetDetailUsersByID(c echo.Context) error {
	employeeID := c.Param("employeeID")

	user, err := ctrl.service.GetDetailUserByID(employeeID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.FromDomainUser(user)

	return c.JSON(http.StatusOK, r.SuccessResponse(result))
}

func (ctrl *Controller) GetAllUsers(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, r.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	users, err := ctrl.service.GetAllUsers(companyID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewGetAllUsersReponse(users)

	return c.JSON(http.StatusOK, r.SuccessResponse(result))
}

func (ctrl *Controller) GetDetailUserDashboard(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	userID := c.Param("employeeID")

	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, r.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	result, err := ctrl.service.GetDetailUserDashboard(userID)
	if err != nil {
		if err == exception.ErrDataNotFound {
			return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	resp := response.NewGetAllUserDetailDashboardResp(result.User, result.Courses)

	return c.JSON(http.StatusOK, r.SuccessResponse(resp))
}

func (ctrl *Controller) UpdateSpecializationName(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	userID := c.Param("employeeID")

	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, r.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	updateSpecializationNameRequest := new(request.UpdateSpecializationNameRequest)
	if err := c.Bind(&updateSpecializationNameRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateSpecializationNameRequest.CompanyID = companyID
	updateSpecializationNameRequest.UserID = userID

	req := *updateSpecializationNameRequest.ToSpec()

	id, err := ctrl.service.UpdateSpecializationName(req)
	if err != nil {
		if err == exception.ErrEmployeeNotFound {
			return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
	}

	result := response.NewUpdateSpecializationNameResponse(id)

	return c.JSON(http.StatusOK, r.SuccessResponse(result))
}

func (ctrl *Controller) UpdateProfile(c echo.Context) error {
	extract, _ := m.ExtractToken(c)
	companyID := c.Param("companyID")
	userID := c.Param("employeeID")

	if companyID != extract.CompanyId {
		return c.JSON(http.StatusUnauthorized, r.UnauthorizedResponse("You are not authorized to access this resource"))
	}

	updateProfileRequest := new(request.UpdateRequestUser)

	// Check headers Content-Type
	if c.Request().Header.Get("Content-Type") != "application/json" {
		if err := c.Bind(updateProfileRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		file, err := c.FormFile("avatar")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		imageURL, err := s3.UploadFileHelper(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
		}

		updateProfileRequest.CompanyID = companyID
		updateProfileRequest.ID = userID
		updateProfileRequest.Avatar = imageURL

		req := *updateProfileRequest.ToSpecUpdateUsers()

		id, err := ctrl.service.UpdateProfile(req)
		if err != nil {
			if err == exception.ErrEmployeeNotFound {
				return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
			}
			return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
		}

		result := response.NewCreateUpdateUserResponse(id)

		return c.JSON(http.StatusOK, r.SuccessResponse(result))
	} else {
		if err := c.Bind(&updateProfileRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updateProfileRequest.CompanyID = companyID
		updateProfileRequest.ID = userID

		req := *updateProfileRequest.ToSpecUpdateUsers()

		id, err := ctrl.service.UpdateProfile(req)
		if err != nil {
			if err == exception.ErrEmployeeNotFound {
				return c.JSON(http.StatusNotFound, r.NotFoundResponse(err.Error()))
			}
			return c.JSON(http.StatusInternalServerError, r.InternalServerErrorResponse(err.Error()))
		}

		result := response.NewCreateUpdateUserResponse(id)

		return c.JSON(http.StatusOK, r.SuccessResponse(result))
	}
}
