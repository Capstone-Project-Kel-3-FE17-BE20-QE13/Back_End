package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/application"
	"JobHuntz/utils/responses"
	"net/http"

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
	newApply := new(ApplyRequest)
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	newApply.JobseekerID = middlewares.ExtractTokenUserId(c)
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newApply)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to CoreProject
	input := MapApplyReqToCoreApply(*newApply)
	_, err = h.applyService.CreateApplication(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create project", nil))
}

func (h *ApplyHandler) GetAllApplications(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.applyService.GetAllApplications(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var storeResponse []ApplyResponse
	for _, v := range result {
		storeResponse = append(storeResponse, MapCoreApplyToApplyRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", storeResponse))
}
