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
	userID := middlewares.ExtractTokenUserId(c)

	myData, _ := h.applyService.GetMyData(userID)

	vacancyID := c.QueryParam("vacancy_id")

	vacancyID_int, err := strconv.Atoi(vacancyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
	}

	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	resCount, _ := h.applyService.CountApplication(dbRaw, userID)

	result, errGet := h.applyService.GetDataCompany(dbRaw, uint(vacancyID_int))
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error get data company", nil))
	}

	fmt.Println("isi data company: ", result)

	newApply := new(ApplyRequest)
	newApply.JobseekerID = userID
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
	err = h.applyService.CreateApplication(input, resCount, myData.Status_Verification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create application", nil))
}

func (h *ApplyHandler) AppHistoryJobseeker(c echo.Context) error {
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

func (h *ApplyHandler) AppHistoryCompany(c echo.Context) error {
	//userID := middlewares.ExtractTokenUserId(c)

	vacancyID := c.QueryParam("vacancy_id")
	vacancyID_int, errConv := strconv.Atoi(vacancyID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, err := h.applyService.GetAllApplicationsCompany(uint(vacancyID_int))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	fmt.Println("isi result: ", result)

	var applyResponse []ApplyResponse
	for _, v := range result {
		applyResponse = append(applyResponse, MapCoreApplyToApplyRes(v))
	}

	fmt.Println("isi applications: ", applyResponse)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "successfully get all applications history", applyResponse))
	// vacancyID := c.QueryParam("vacancy_id")
	// vacancyID_int, errConv := strconv.Atoi(vacancyID)
	// if errConv != nil {
	// 	return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	// }

	// cfg := config.InitConfig()
	// dbRaw := database.InitRawSql(cfg)

	// result, errFirst := handler.applyService.GetAllApplicationsCompany(dbRaw, vacancyID_int)
	// if errFirst != nil {
	// 	return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	// }

	// return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data applicants", result))
}

func (h *ApplyHandler) EditApplicationStatus(c echo.Context) error {
	applicationId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "invalid user ID", nil))
	}

	editApplication := ApplicationRequestStatus{}
	errBind := c.Bind(&editApplication)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error bind data. data not valid"+errBind.Error(), nil))
	}

	applicationCore := RequestToCore(editApplication)

	errEdit := h.applyService.EditApplication(uint(applicationId), applicationCore)
	if errEdit != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error editing data"+errBind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create application", nil))
}
