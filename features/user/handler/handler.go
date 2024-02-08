package handler

import (
	"JobHuntz/features/user"
	"JobHuntz/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	userCore := RequestUserToCore(newUser)

	errCreate := handler.userService.Register(userCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	userResponse := CoreUserToResponse(userCore)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully registered", userResponse))
}

func (handler *UserHandler) CreateCareer(c echo.Context) error {
	userID := 1

	newCareer := CareerRequest{}
	newCareer.UserID = uint(userID)

	errBind := c.Bind(&newCareer)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	careerCore := RequestCareerToCore(newCareer)

	errCreate := handler.userService.AddCareer(careerCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	careerResponse := CoreCareerToResponse(careerCore)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create career", careerResponse))
}
