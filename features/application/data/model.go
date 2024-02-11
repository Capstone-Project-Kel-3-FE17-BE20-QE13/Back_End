package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"
)

func CoreToModel(input application.Core) database.Application {
	return database.Application{
		// ID:          input.Id,
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Status_application: input.Status_application,
	}
}

func CoretoModelGorm(data []application.Core) []database.Application {
	var productsDataGorm []database.Application
	for _, input := range data {
		var productGorm = database.Application{
			JobseekerID:        input.JobseekerID,
			VacancyID:          input.VacancyID,
			Status_application: input.Status_application,
		}
		productsDataGorm = append(productsDataGorm, productGorm)
	}

	return productsDataGorm
}

func ModelToCore(input database.Application) application.Core {
	return application.Core{
		ID:                 input.ID,
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Status_application: input.Status_application,
	}
}

func ModelGormToCore(data []database.Application) []application.Core {
	var productsData []application.Core
	for _, input := range data {
		var productInput = application.Core{
			JobseekerID:        input.JobseekerID,
			VacancyID:          input.VacancyID,
			Status_application: input.Status_application,
		}
		productsData = append(productsData, productInput)
	}

	return productsData
}
