package handler

import (
	"JobHuntz/features/payment"
	"JobHuntz/utils"
)

type paymentResponse struct {
	ID          string          `json:"id"`
	OrderID     string          `json:"order_id"`
	Amount      string          `json:"amount"`
	BankAccount string          `json:"bank_account"`
	VANumber    string          `json:"va_number"`
	UserID      uint            `json:"user_id" form:"user_id"`
	Status      string          `json:"status"`
	CreatedAt   utils.LocalTime `json:"created_at"`
	UpdatedAt   utils.LocalTime `json:"updated_at"`
}

func paymentResp(p payment.PaymentCore) paymentResponse {
	return paymentResponse{
		ID:          p.ID,
		OrderID:     p.OrderID,
		Amount:      p.Amount,
		BankAccount: p.BankAccount,
		VANumber:    p.VANumber,
		UserID:      p.UserID,
		Status:      p.Status,
		CreatedAt:   utils.LocalTime(p.CreatedAt),
		UpdatedAt:   utils.LocalTime(p.UpdatedAt),
	}
}
