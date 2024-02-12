package verification

type OrderJobseekerCore struct {
	ID           uint    `json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

type OrderCompanyCore struct {
	ID           uint    `json:"id" form:"id"`
	CompanyID    uint    `json:"company_id" form:"company_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

// interface untuk Service Layer
type VerificationServiceInterface interface {
	AddOrderJobseeker(input OrderJobseekerCore) error
	AddOrderCompany(input OrderCompanyCore) error
}

// interface untuk Data Layer
type VerificationDataInterface interface {
	AddOrderJobseeker(input OrderJobseekerCore) error
	AddOrderCompany(input OrderCompanyCore) error
}
