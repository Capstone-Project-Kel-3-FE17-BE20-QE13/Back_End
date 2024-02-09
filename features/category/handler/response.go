package handler

import (
	"JobHuntz/features/category"
)

type CategoryResponse struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

func FromDomain(domain category.Core) CategoryResponse {
	return CategoryResponse{
		Id:       domain.Id,
		Category: domain.Category,
	}
}

func ListFromDomain(domain []category.Core) (response []CategoryResponse) {
	for _, Category := range domain {
		response = append(response, FromDomain(Category))
	}
	return
}
