package handler

import (
	"JobHuntz/app/middlewares"
	jobs "JobHuntz/features/vacancy"
	"JobHuntz/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JobHandler struct {
	jobService jobs.JobServiceInterface
}

func NewJob(jobService jobs.JobServiceInterface) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

// insert product
func (handler *JobHandler) CreateJobs(c echo.Context) error {
	newJob := new(JobRequest)
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	newJob.CompanyID = middlewares.ExtractTokenUserId(c)

	errBind := c.Bind(&newJob)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	jobCore := RequestToCore(newJob)

	errCreate := handler.jobService.CreateJob(jobCore) // CreateJob(jobCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create job", nil))
}

// delete product
func (handler *JobHandler) Delete(c echo.Context) error {
	jobID := c.Param("job_id")

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

// get all products
func (handler *JobHandler) GetAllJob(c echo.Context) error {
	result, errFind := handler.jobService.GetAllJobs()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

func (handler *JobHandler) GetJobById(c echo.Context) error {
	jobId := c.Param("job_id")

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
