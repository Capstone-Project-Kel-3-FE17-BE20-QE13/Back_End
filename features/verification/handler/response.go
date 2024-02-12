package handler

type OrderResponse struct {
	ID           string  `gorm:"type:varchar(40);primary_key" json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}
