package handler

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/verification"
	"JobHuntz/utils/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VerificationHandler struct {
	verificationService verification.VerificationServiceInterface
}

func New(service verification.VerificationServiceInterface) *VerificationHandler {
	return &VerificationHandler{
		verificationService: service,
	}
}

// insert order
func (handler *VerificationHandler) CreateOrderJobseeker(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	newOrder := OrderJobseekerRequest{}
	newOrder.JobseekerID = userID

	fmt.Println("isi user id: ", userID)

	newOrder.Price = 1000000
	newOrder.Status_order = "On Going"

	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
	}

	fmt.Println("isi order: ", newOrder)

	orderCore := RequestOrderJobseekerToCore(newOrder)

	errCreate := handler.verificationService.AddOrderJobseeker(orderCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create order", nil))
}

func (handler *VerificationHandler) CreateOrderCompany(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	newOrder := OrderCompanyRequest{}
	newOrder.CompanyID = userID

	fmt.Println("isi user id: ", &userID)

	newOrder.Price = 2000000
	newOrder.Status_order = "On Going"

	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
	}

	orderCore := RequestOrderCompanyToCore(newOrder)

	errCreate := handler.verificationService.AddOrderCompany(orderCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create order", nil))
}
