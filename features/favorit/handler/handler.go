package handler

import (
	config "JobHuntz/app/configs"
	"JobHuntz/app/database"
	"JobHuntz/app/middlewares"
	"JobHuntz/features/favorit"
	"JobHuntz/utils/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FavHandler struct {
	favService favorit.FavServiceInterface
}

func New(service favorit.FavServiceInterface) *FavHandler {
	return &FavHandler{
		favService: service,
	}
}

func (h *FavHandler) CreateFavorit(c echo.Context) error {
	vacancyID := c.QueryParam("vacancy_id")

	vacancyID_int, err := strconv.Atoi(vacancyID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
	}

	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	result, errGet := h.favService.GetDataCompany(dbRaw, uint(vacancyID_int))
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error get data company", nil))
	}

	fmt.Println("isi data company: ", result)

	newFavor := new(FavRequest)
	newFavor.JobseekerID = middlewares.ExtractTokenUserId(c)
	newFavor.VacancyID = uint(vacancyID_int)
	newFavor.Position = result.Position
	newFavor.Company_name = result.Company_name

	errFavor := c.Bind(&newFavor)
	if errFavor != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	fmt.Println("isi new favor: ", newFavor)

	//mapping dari request to Core
	input := MapApplyReqToCoreApply(*newFavor)
	_, err = h.favService.CreateFavorit(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create favourite", nil))
}

func (h *FavHandler) GetAllFavorit(c echo.Context) error {
	// all favourites didapatkan berdasarkan jobseeker id yg login
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.favService.GetAllFavorit(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var storeResponse []FavResponse
	for _, v := range result {
		storeResponse = append(storeResponse, MapCoreApplyToApplyRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", storeResponse))
}

func (h *FavHandler) DeleteFavById(c echo.Context) error {
	favID := c.Param("favorit_id")
	userID := middlewares.ExtractTokenUserId(c)

	FavID_int, errConv := strconv.Atoi(favID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errRead := h.favService.GetAllFavorit(userID)
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	errDel := h.favService.DeleteFavById(result, FavID_int)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete favourite", nil))
}

// func (h *FavHandler) DeleteFavById(c echo.Context) error {
// 	userID := middlewares.ExtractTokenUserId(c)
// 	err := h.favService.DeleteFavById(uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete user", nil))
// }
