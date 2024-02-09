package handler

import "JobHuntz/features/category"

type CategoryCreate struct {
	Category string `json:"category"`
}

func ToDomain(Ctg *CategoryCreate) category.Core {
	return category.Core{
		Category: Ctg.Category,
	}
}
