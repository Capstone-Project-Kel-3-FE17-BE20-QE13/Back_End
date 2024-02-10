package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"fmt"
	"net/http"
	"time"

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

func (handler *JobseekerHandler) LoginJobseeker(c echo.Context) error {
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

func (handler *JobseekerHandler) UpdateJobseeker(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newUpdate := JobseekerRequest{}

	birthDateString := c.FormValue("birth_date")
	if birthDateString != "" {
		birthDate, err := time.Parse("2006-01-02", birthDateString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		newUpdate.Birth_date = birthDate
	}

	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println(newUpdate)

	newUpdateCore := RequestJobseekerToCore(newUpdate)

	err := handler.jobseekerService.UpdateProfile(seekerID, newUpdateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update profile", nil))
}

func (handler *JobseekerHandler) CreateCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	inputCV, errRead := c.FormFile("cv_file")
	if errRead != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot read file", nil))
	}

	responURL, errURL := handler.jobseekerService.CV(inputCV)
	if errURL != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot get resp", nil))
	}

	newCV := CVRequest{}
	newCV.JobseekerID = seekerID
	newCV.CV_file = responURL.SecureURL

	errBind := c.Bind(&newCV)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	cvCore := RequestCVToCore(newCV)

	errCreate := handler.jobseekerService.AddCV(cvCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully upload cv", nil))
}

func (handler *JobseekerHandler) GetCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errFirst := handler.jobseekerService.ReadCV(seekerID)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get cv", result))
}

func (handler *JobseekerHandler) UpdateCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newCV := CVRequest{}
	newCV.JobseekerID = seekerID
	oldCV, _ := c.FormFile("cv_file")
	if oldCV != nil {
		responURL, errResp := handler.jobseekerService.CV(oldCV)
		if errResp != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get resp "+errResp.Error(), nil))
		}
		newCV.CV_file = responURL.SecureURL
	}

	errBind := c.Bind(&newCV)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	newCVCore := RequestCVToCore(newCV)

	fmt.Println("isi update: ", newCVCore)

	errUpdate := handler.jobseekerService.UpdateCV(newCVCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update cv", nil))
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
