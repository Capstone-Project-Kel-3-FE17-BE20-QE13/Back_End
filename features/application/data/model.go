package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"
)

func CoreToModel(input application.Core) database.Apply {
	return database.Apply{
		// ID:          input.Id,
		JobseekerID: input.JobseekerID,
		JobId:       input.JobId,
		Status:      input.Status,
	}
}

func CoretoModelGorm(data []application.Core) []database.Apply {
	var productsDataGorm []database.Apply
	for _, input := range data {
		var productGorm = database.Apply{
			JobseekerID: input.JobseekerID,
			JobId:       input.JobId,
			Status:      input.Status,
		}
		productsDataGorm = append(productsDataGorm, productGorm)
	}

	return productsDataGorm
}

func ModelToCore(input database.Apply) application.Core {
	return application.Core{
		JobseekerID: input.JobseekerID,
		JobId:       input.JobId,
		Status:      input.Status,
	}
}

func ModelGormToCore(data []database.Apply) []application.Core {
	var productsData []application.Core
	for _, input := range data {
		var productInput = application.Core{
			JobseekerID: input.JobseekerID,
			JobId:       input.JobId,
			Status:      input.Status,
		}
		productsData = append(productsData, productInput)
	}

	return productsData
}
