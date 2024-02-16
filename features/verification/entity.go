package verification

import (
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
)

type OrderCore struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	JobseekerID  *uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CompanyID    *uint   `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderJobseekerCore struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderCompanyCore struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	CompanyID    uint    `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

// interface untuk Service Layer
type VerificationServiceInterface interface {
	GetDataJobseeker(userID uint) (jobseeker.JobseekerCore, error)
	GetDataCompany(userID uint) (company.CompanyCore, error)
	AddOrderJobseeker(input OrderJobseekerCore) error
	AddOrderCompany(input OrderCompanyCore) error
}

// interface untuk Data Layer
type VerificationDataInterface interface {
	GetDataJobseeker(userID uint) (jobseeker.JobseekerCore, error)
	GetDataCompany(userID uint) (company.CompanyCore, error)
	AddOrderJobseeker(input OrderJobseekerCore) error
	AddOrderCompany(input OrderCompanyCore) error
}
