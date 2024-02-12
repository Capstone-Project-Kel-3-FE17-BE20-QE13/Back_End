package verification

type OrderCore struct {
	ID           uint    `json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

// interface untuk Service Layer
type VerificationServiceInterface interface {
	AddOrder(input OrderCore) error
}

// interface untuk Data Layer
type VerificationDataInterface interface {
	AddOrder(input OrderCore) error
}
