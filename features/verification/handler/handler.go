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
func (handler *VerificationHandler) CreateOrder(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	seekerData, _ := handler.verificationService.GetDataJobseeker(userID)
	fmt.Println("isi seeker: ", seekerData)

	if seekerData.ID != 0 {
		newOrder := OrderRequest{}
		newOrder.JobseekerID = &userID

		fmt.Println("isi pointer user id: ", &userID)

		newOrder.CompanyID = nil
		newOrder.Price = 1000000
		newOrder.Status_order = "On Going"

		errBind := c.Bind(&newOrder)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
		}

		orderCore := RequestOrderToCore(newOrder)

		errCreate := handler.verificationService.AddOrder(orderCore)
		if errCreate != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
		}
	} else if seekerData.ID == 0 {
		companyData, _ := handler.verificationService.GetDataCompany(userID)
		if companyData.ID != 0 {
			newOrder := OrderRequest{}
			newOrder.JobseekerID = nil
			newOrder.CompanyID = &userID
			newOrder.Price = 2000000
			newOrder.Status_order = "On Going"

			errBind := c.Bind(&newOrder)
			if errBind != nil {
				return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
			}

			orderCore := RequestOrderToCore(newOrder)

			errCreate := handler.verificationService.AddOrder(orderCore)
			if errCreate != nil {
				return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
			}
		}
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success create order", nil))
}
