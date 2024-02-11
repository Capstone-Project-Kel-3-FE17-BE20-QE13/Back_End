package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/favorit"
	"JobHuntz/utils/responses"
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
	newApply := new(FavRequest)
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	newApply.JobId = middlewares.ExtractTokenUserId(c)
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newApply)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to CoreProject
	input := MapApplyReqToCoreApply(*newApply)
	_, err = h.favService.CreateFavorit(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create project", nil))
}

// func (h *FavHandler) GetAllFavorit(c echo.Context) error {
// 	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
// 	userID := middlewares.ExtractTokenUserId(c)
// 	result, err := h.favService.GetAllFavorit(userID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
// 	}
// 	var storeResponse []FavResponse
// 	for _, v := range result {
// 		storeResponse = append(storeResponse, MapCoreApplyToApplyRes(v))
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", storeResponse))
// }

func (h *FavHandler) GetAllFavorit(c echo.Context) error {
	result, errFind := h.favService.GetAllFavorit()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}
func (h *FavHandler) DeleteFavById(c echo.Context) error {
	favID := c.Param("fav_id")

	FavID_int, errConv := strconv.Atoi(favID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errRead := h.favService.GetAllFavorit()
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	errDel := h.favService.DeleteFavById(result, FavID_int)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete data", nil))
}

// func (h *FavHandler) DeleteFavById(c echo.Context) error {
// 	userID := middlewares.ExtractTokenUserId(c)
// 	err := h.favService.DeleteFavById(uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete user", nil))
// }
