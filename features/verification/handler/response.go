package handler

type OrderResponse struct {
	ID           uint    `json:"id" form:"id"`
	JobseekerID  uint    `json:"jobseeker_id" form:"jobseeker_id"`
	Price        float64 `json:"price" form:"price"`
	Status_order string  `json:"stat_order" form:"stat_order"`
}

// func CoreToResponse(input order.Core) OrderResponse {
// 	return OrderResponse{
// 		ID:             input.ID,
// 		ShoppingCartID: input.ShoppingCartID,
// 		Item:           input.Item,
// 		Status:         input.Status,
// 		CreatedAt:      input.CreatedAt,
// 		UpdatedAt:      input.UpdatedAt,
// 	}
// }

// func CoreToResponseList(data []order.Core) []OrderResponse {
// 	var results []OrderResponse
// 	for _, v := range data {
// 		results = append(results, CoreToResponse(v))
// 	}
// 	return results
// }
