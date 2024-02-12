package handler

import (
	config "JobHuntz/app/configs"
	"JobHuntz/app/database"
	"JobHuntz/app/middlewares"
	"JobHuntz/features/application"
	"JobHuntz/utils/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ApplyHandler struct {
	applyService application.ApplyServiceInterface
}

func New(service application.ApplyServiceInterface) *ApplyHandler {
	return &ApplyHandler{
		applyService: service,
	}
}

func (h *ApplyHandler) CreateApply(c echo.Context) error {
	vacancyID := c.QueryParam("vacancy_id")

	vacancyID_int, err := strconv.Atoi(vacancyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
	}

	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	result, errGet := h.applyService.GetDataCompany(dbRaw, uint(vacancyID_int))
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error get data company", nil))
	}

	fmt.Println("isi data company: ", result)

	newApply := new(ApplyRequest)
	newApply.JobseekerID = middlewares.ExtractTokenUserId(c)
	newApply.VacancyID = uint(vacancyID_int)
	newApply.Position = result.Position
	newApply.Company_name = result.Company_name
	newApply.Status_application = "Dikirim"

	errBind := c.Bind(&newApply)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to Core
	input := MapApplyReqToCoreApply(*newApply)
	_, err = h.applyService.CreateApplication(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create application", nil))
}

func (h *ApplyHandler) GetApplicationsHistory(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.applyService.GetAllApplications(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var applyResponse []ApplyResponse
	for _, v := range result {
		applyResponse = append(applyResponse, MapCoreApplyToApplyRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "successfully get all applications history", applyResponse))
}
