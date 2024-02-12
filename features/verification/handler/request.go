package handler

import (
	"JobHuntz/features/verification"
)

type OrderRequest struct {
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

func RequestOrderToCore(input OrderRequest) verification.OrderCore {
	return verification.OrderCore{
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
