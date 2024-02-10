package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"

	"gorm.io/gorm"
)

type ApplyQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) application.ApplyDataInterface {
	return &ApplyQuery{
		db: db,
	}
}

func (repo *ApplyQuery) CreateApplication(input application.Core) (uint, error) {
	// simpan ke DB
	newProductGorm := CoreToModel(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return 0, tx.Error
	}

	return newProductGorm.ID, nil
}

func (repo *ApplyQuery) GetAllApplications(jobseekerID uint) ([]application.Core, error) {
	var productsDataGorm []database.Apply
	tx := repo.db.Find(&productsDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProductCore := ModelGormToCore(productsDataGorm)

	return allProductCore, nil
}
