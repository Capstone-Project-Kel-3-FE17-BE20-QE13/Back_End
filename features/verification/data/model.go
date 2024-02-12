package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/verification"
)

func CoreJobseekerToModel(input verification.OrderJobseekerCore) database.Order {
	return database.Order{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}

func CoreCompanyToModel(input verification.OrderCompanyCore) database.OrderCompany {
	return database.OrderCompany{
		CompanyID:    input.CompanyID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
