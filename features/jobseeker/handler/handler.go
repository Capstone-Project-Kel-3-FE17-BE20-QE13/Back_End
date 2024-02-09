package handler

import (
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type JobseekerHandler struct {
	jobseekerService jobseeker.JobseekerServiceInterface
}

func New(service jobseeker.JobseekerServiceInterface) *JobseekerHandler {
	return &JobseekerHandler{
		jobseekerService: service,
	}
}

func (handler *JobseekerHandler) RegisterJobseeker(c echo.Context) error {
	newSeeker := JobseekerRequest{}
	errBind := c.Bind(&newSeeker)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	seekerCore := RequestJobseekerToCore(newSeeker)

	errCreate := handler.jobseekerService.Register(seekerCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully registered", nil))
}

func (handler *JobseekerHandler) Login(c echo.Context) error {
	newLogin := JobseekerRequest{}
	errBind := c.Bind(&newLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	resLogin, token, err := handler.jobseekerService.Login(newLogin.Email, newLogin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	loginResponse := CoreJobseekerToResponseLogin(resLogin, token)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully login", loginResponse))
}

// func (handler *UserHandler) CreateCareer(c echo.Context) error {
// 	seekerID := 1

// 	newCareer := CareerRequest{}
// 	newCareer.JobseekerID = uint(seekerID)

// 	errBind := c.Bind(&newCareer)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
// 	}

// 	careerCore := RequestCareerToCore(newCareer)

// 	errCreate := handler.userService.AddCareer(careerCore)
// 	if errCreate != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
// 	}

// 	careerResponse := CoreCareerToResponse(careerCore)

// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create career", careerResponse))
// }
