package handler

import (
	"JobHuntz/features/company"
	"JobHuntz/utils/responses"
	"net/http"

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
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	companyCore := RequestCompanyToCore(newCompany)

	errCreate := handler.companyService.RegisterCompany(companyCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully registered", nil))
}

func (handler *CompanyHandler) LoginCompany(c echo.Context) error {
	companyLogin := CompanyRequestLogin{}
	errBind := c.Bind(&companyLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	ressLogin, token, err := handler.companyService.LoginCompany(companyLogin.Email, companyLogin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	loginResponse := ResponCompanyToResponseLogin(ressLogin, token)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully login", loginResponse))
}
