package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/favorit"
)

func CoreToModel(input favorit.Core) database.Favorites {
	return database.Favorites{
		// ID:          input.Id,
		JobseekerID: input.JobseekerID,
		JobId:       input.JobId,
		Name:        input.Name,
	}
}

func CoretoModelGorm(data []favorit.Core) []database.Favorites {
	var productsDataGorm []database.Favorites
	for _, input := range data {
		var productGorm = database.Favorites{
			JobseekerID: input.JobseekerID,
			JobId:       input.JobId,
			Name:        input.Name,
		}
		productsDataGorm = append(productsDataGorm, productGorm)
	}

	return productsDataGorm
}

func ModelToCore(input database.Favorites) favorit.Core {
	return favorit.Core{
		JobseekerID: input.JobseekerID,
		JobId:       input.JobId,
		Name:        input.Name,
	}
}

func ModelGormToCore(data []database.Favorites) []favorit.Core {
	var productsData []favorit.Core
	for _, input := range data {
		var productInput = favorit.Core{
			JobseekerID: input.JobseekerID,
			JobId:       input.JobId,
			Name:        input.Name,
		}
		productsData = append(productsData, productInput)
	}

	return productsData
}
