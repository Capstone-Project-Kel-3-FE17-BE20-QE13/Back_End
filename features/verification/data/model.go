package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/verification"
)

func HistoryToModel(input verification.OrderCore) database.OrderJobseeker {
	return database.OrderJobseeker{
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
