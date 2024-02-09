package handler

import (
	"JobHuntz/features/category"
	"JobHuntz/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Ct category.CategoryDataInterface
}

func NewCategory(Ct category.CategoryDataInterface) *CategoryHandler {
	return &CategoryHandler{
		Ct: Ct,
	}
}

// func (CategoryHandler CategoryHandler) CreateCategory(c echo.Context) error {
// 	categoryCreate := new(CategoryCreate)
// 	errBind := c.Bind(&categoryCreate)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
// 	}
// 	CategoryCore := ToDomain(categoryCreate)

// 	errCreate := CategoryHandler.Ct.CreateCategory(CategoryCore) // CreateJob(jobCore)
// 	if errCreate != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create category", nil))
// }

func (CategoryHandler *CategoryHandler) GetAllCategory(c echo.Context) error {
	// fmt.Println("GetAllCategories")
	result, errFind := CategoryHandler.Ct.GetAllCategory()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

func (CategoryHandler CategoryHandler) GetCategoryById(c echo.Context) error {
	categoryID := c.Param("catergory_id")

	ID, errConv := strconv.Atoi(categoryID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := CategoryHandler.Ct.GetCategoryById(ID)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

// func (CategoryController CategoryController) DeleteAllcategories(c echo.Context) error {
// 	fmt.Println("HardDeleteAllCategoryRecords")

// 	ctx := c.Request().Context()
// 	error := CategoryController.CategoryUseCase.DeleteAllcategories(ctx)

// 	if error != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(categories.Domain{}))
// }

// func (CategoryHandler CategoryHandler) DeleteCategoryById(c echo.Context) error {
// 	fmt.Println("DeleteCategoryById")

// 	categoryId, err := strconv.Atoi(c.Param("categoryId"))
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	ctx := c.Request().Context()
// 	category, err := CategoryHandler.Ct.DeleteCategoryById(ctx, categoryId)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(category))
// }
