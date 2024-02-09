package service

import (
	"JobHuntz/features/category"
)

type CategoryService struct {
	Repo category.CategoryServiceInterface
}

func NewCategory(repo category.CategoryServiceInterface) category.CategoryDataInterface {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) GetAllCategory() ([]category.Core, error) {
	result, err := s.Repo.GetAllCategory()

	return result, err
}

func (s *CategoryService) GetCategoryById(id int) (category.Core, error) {
	result, err := s.Repo.GetCategoryById(id)
	return result, err
}
