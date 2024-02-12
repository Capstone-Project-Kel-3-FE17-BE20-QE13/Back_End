package handler

import (
	"JobHuntz/features/verification"
)

type OrderJobseekerRequest struct {
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderCompanyRequest struct {
	CompanyID    uint    `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

func RequestOrderJobseekerToCore(input OrderJobseekerRequest) verification.OrderJobseekerCore {
	return verification.OrderJobseekerCore{
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}

func RequestOrderCompanyToCore(input OrderCompanyRequest) verification.OrderCompanyCore {
	return verification.OrderCompanyCore{
		CompanyID:    input.CompanyID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
