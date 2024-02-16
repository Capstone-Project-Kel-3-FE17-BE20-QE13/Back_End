package handler

import (
	"JobHuntz/features/verification"

	"github.com/google/uuid"
)

type OrderRequest struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	JobseekerID  *uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CompanyID    *uint   `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderJobseekerRequest struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderCompanyRequest struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	CompanyID    uint    `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

func RequestOrderToCore(input OrderRequest) verification.OrderCore {
	return verification.OrderCore{
		ID:           uuid.New().String(),
		JobseekerID:  input.JobseekerID,
		CompanyID:    input.CompanyID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}

func RequestOrderJobseekerToCore(input OrderJobseekerRequest) verification.OrderJobseekerCore {
	return verification.OrderJobseekerCore{
		ID:           uuid.New().String(),
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}

func RequestOrderCompanyToCore(input OrderCompanyRequest) verification.OrderCompanyCore {
	return verification.OrderCompanyCore{
		ID:           uuid.New().String(),
		CompanyID:    input.CompanyID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
