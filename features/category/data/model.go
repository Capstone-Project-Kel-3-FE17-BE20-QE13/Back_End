package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/category"
)

func CoreToModel(input category.Core) database.Categories {
	return database.Categories{
		Id:       input.Id,
		Category: input.Category,
	}
}

func CoretoModelGorm(data []category.Core) []database.Categories {
	var CategoriesDataGorm []database.Categories
	for _, input := range data {
		var CategoriesGorm = database.Categories{
			Id:       input.Id,
			Category: input.Category,
		}
		CategoriesDataGorm = append(CategoriesDataGorm, CategoriesGorm)
	}
	return CategoriesDataGorm
}

func ModelToCore(input database.Categories) category.Core {
	return category.Core{
		Id:       input.Id,
		Category: input.Category,
	}
}

func ModelGormToCore(data []database.Categories) []category.Core {
	var CategoriesData []category.Core
	for _, input := range data {
		var CategoriesInput = category.Core{
			Id:       input.Id,
			Category: input.Category,
		}
		CategoriesData = append(CategoriesData, CategoriesInput)
	}

	return CategoriesData
}
