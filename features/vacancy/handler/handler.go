package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/vacancy"
	"JobHuntz/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JobHandler struct {
	jobService vacancy.JobServiceInterface
}

func NewJob(jobService vacancy.JobServiceInterface) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

func (handler *JobHandler) GetAllJob(c echo.Context) error {
	result, errFind := handler.jobService.GetAllJobs()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

func (handler *JobHandler) GetJobById(c echo.Context) error {
	jobId := c.Param("vacancy_id")

	jobId_int, errConv := strconv.Atoi(jobId)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobService.GetJobById(jobId_int)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

func (handler *JobHandler) CreateJobs(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, _ := handler.jobService.GetById(userId)
	namaUser := user.Status_Verification

	if namaUser == "Unverified" {
		jobCount, _ := handler.jobService.CountJobsByUserID(userId)
		if jobCount >= 3 {
			return c.JSON(http.StatusForbidden, responses.WebResponse(http.StatusForbidden, "Unverified users can only create 3 jobs", nil))
		}
	}

	newJob := JobRequest{}
	errBind := c.Bind(&newJob)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	jobCore := vacancy.Core{
		CompanyID:       userId,
		Name:            newJob.Name,
		Job_type:        newJob.Job_type,
		Salary_range:    newJob.Salary_range,
		Category:        newJob.Category,
		Job_description: newJob.Job_description,
		Job_requirement: newJob.Job_requirement,
		Status:          "Dibuka",
	}

	errCreate := handler.jobService.CreateJob(jobCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create job", nil))
}

func (handler *JobHandler) GetVacanciesMadeByCompany(c echo.Context) error {
	companyID := middlewares.ExtractTokenUserId(c)

	result, errFind := handler.jobService.MyCompanyVacancies(companyID)
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

func (handler *JobHandler) Delete(c echo.Context) error {
	jobID := c.Param("vacancy_id")

	ID, errConv := strconv.Atoi(jobID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errRead := handler.jobService.GetAllJobs()
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	errDel := handler.jobService.DeleteJobById(result, ID)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *JobHandler) UpdateVacancyStatus(c echo.Context) error {
	vacancyID := c.Param("vacancy_id")

	vacancyID_int, errConv := strconv.Atoi(vacancyID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newStatus := JobStatusRequest{}
	errBind := c.Bind(&newStatus)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	newStatusCore := RequestStatusToCore(newStatus)

	errUpdate := handler.jobService.UpdateStatus(newStatusCore, uint(vacancyID_int))
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success update status", nil))
}
