package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/category"

	"gorm.io/gorm"
)

type CategoryQuery struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) category.CategoryDataInterface {
	return &CategoryQuery{
		db: db,
	}
}

func (repo *CategoryQuery) GetCategoryById(id int) (category.Core, error) {
	var Category database.Categories
	result := repo.db.First(&Category, id)

	if result.Error != nil {
		return category.Core{}, result.Error
	}

	singleCategory := ModelToCore(Category)

	return singleCategory, nil
}

func (repo *CategoryQuery) GetAllCategory() ([]category.Core, error) {
	var newCatagoryGorm []database.Categories
	tx := repo.db.Find(&newCatagoryGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allCatagoryCore := ModelGormToCore(newCatagoryGorm)

	return allCatagoryCore, nil
}

// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
// }
// func (rep *MysqlCategoryRepository) UpdateCat(ctx context.Context, domain categories.Domain) (Cats.Domain, error) {
// 	var category Categories
// 	result := rep.Conn.First(&category, "email = ? AND password = ?", domain.Email, domain.Password)

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Model(&Cat).Updates(FromDomain(domain))

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	result = rep.Conn.Save(&Cat)

// 	if result.Error != nil {
// 		return Cats.Domain{}, result.Error
// 	}

// 	return Cat.ToDomain(), nil
// }

// func (repo *CategoryQuery) DeleteCategoryById(ctx context.Context, id int) (category.Core, error) {
// 	var Category Categories
// 	result := repo.Conn.Where("id = ?", id).Delete(&Category)

// 	if result.Error != nil {
// 		return category.Core{}, result.Error
// 	}

// 	return data.ToDomain(), nil
// }
