package category

import (
	"time"
)

type Core struct {
	Id        int
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryDataInterface interface {
	//CreateCategory(input Core) (int, error)
	GetCategoryById(id int) (Core, error)
	GetAllCategory() ([]Core, error)
	//DeleteCategoryById(id int) error
}

type CategoryServiceInterface interface {
	//CreateCategory(input Core) (int, error)
	GetCategoryById(id int) (Core, error)
	GetAllCategory() ([]Core, error)
	//DeleteCategoryById(id int) error
}
