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

func (handler *UserHandler) Login(c echo.Context) error {
	newLogin := UserRequest{}
	errBind := c.Bind(&newLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	resLogin, token, err := handler.userService.Login(newLogin.Email, newLogin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	loginResponse := CoreUserToResponseLogin(resLogin, token)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully login", loginResponse))
}

func (handler *UserHandler) CreateCareer(c echo.Context) error {
	seekerID := 1

	newCareer := CareerRequest{}
	newCareer.JobseekerID = uint(seekerID)

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
