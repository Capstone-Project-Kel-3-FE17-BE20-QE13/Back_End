package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/company"
	"JobHuntz/utils/responses"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CompanyHandler struct {
	companyService company.CompanyServiceInterface
}

func New(service company.CompanyServiceInterface) *CompanyHandler {
	return &CompanyHandler{
		companyService: service,
	}
}

func (handler *CompanyHandler) RegisterCompany(c echo.Context) error {
	newCompany := CompanyRequest{}
	errBind := c.Bind(&newCompany)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error bind data. data not valid", nil))
	}

	companyCore := RequestCompanyToCore(newCompany)

	_, _, errCreate := handler.companyService.RegisterCompany(companyCore)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully registered", nil))
}

func (handler *CompanyHandler) LoginCompany(c echo.Context) error {
	var reqData = CompanyRequestLogin{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	result, token, err := handler.companyService.LoginCompany(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, responses.WebResponse(http.StatusForbidden, "Email atau password tidak boleh kosong ", nil))
	}

	responData := map[string]any{
		"id":        result.ID,
		"full_name": result.Full_name,
		"email":     result.Email,
		"roles":     "company",
		"token":     token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success login.", responData))
}

func (handler *CompanyHandler) GetById(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errGetByID := handler.companyService.GetById(seekerID)
	if errGetByID != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("isi url gambar: ", result.Banners)

	companyresul := CoreResponGetByid(*result)

	fmt.Println("isi url gambar: ", companyresul.Banners)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success.", companyresul))
}

func (handler *CompanyHandler) UpdateCompany(c echo.Context) error {
	var fileSize int64
	var nameFile string

	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "unauthorized or jwt expired", nil))
	}

	var reqData = CompanyRequestUpdate{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	companyCore := RequestCompanyUpdateToCore(reqData)

	fmt.Println("data:", companyCore)

	fileHeader, _ := c.FormFile("banners")

	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error invalid type format, format file not valid", nil))
		}

		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error size data, file size is too big", nil))
		}
	}

	newPhone := c.FormValue("phone")
	if newPhone != "" {
		if _, err := strconv.Atoi(reqData.Phone); err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error invalid phone number format, only numeric allowed", nil))
		}
	}

	err := handler.companyService.UpdateCompany(int(idJWT), companyCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data. update failed", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success.", nil))
}
