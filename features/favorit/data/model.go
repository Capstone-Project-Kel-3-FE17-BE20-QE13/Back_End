package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/favorit"
)

func CoreToModel(input favorit.Core) database.Favourite {
	return database.Favourite{
		// ID:          input.Id,
		JobseekerID:  input.JobseekerID,
		VacancyID:    input.VacancyID,
		Position:     input.Position,
		Company_name: input.Company_name,
	}
}

func CoretoModelGorm(data []favorit.Core) []database.Favourite {
	var productsDataGorm []database.Favourite
	for _, input := range data {
		var productGorm = database.Favourite{
			JobseekerID:  input.JobseekerID,
			VacancyID:    input.VacancyID,
			Position:     input.Position,
			Company_name: input.Company_name,
		}
		productsDataGorm = append(productsDataGorm, productGorm)
	}

	return productsDataGorm
}

func ModelToCore(input database.Favourite) favorit.Core {
	return favorit.Core{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		VacancyID:    input.VacancyID,
		Position:     input.Position,
		Company_name: input.Company_name,
	}
}

func ModelGormToCore(data []database.Favourite) []favorit.Core {
	var productsData []favorit.Core
	for _, input := range data {
		var productInput = favorit.Core{
			ID:           input.ID,
			JobseekerID:  input.JobseekerID,
			VacancyID:    input.VacancyID,
			Position:     input.Position,
			Company_name: input.Company_name,
		}
		productsData = append(productsData, productInput)
	}

	return productsData
}
