package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/favorit"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type FavQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) favorit.FavDataInterface {
	return &FavQuery{
		db: db,
	}
}

func (repo *FavQuery) GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error) {
	// simpan ke DB
	var dataCompany favorit.DataCompanyCore

	query := `SELECT vacancies.name, companies.company_name 
	FROM companies JOIN vacancies ON companies.id = vacancies.company_id
	WHERE vacancies.id = ?`

	tx := dbRaw.QueryRow(query, vacancyID)
	if err := tx.Scan(&dataCompany.Position, &dataCompany.Company_name); err != nil {
		//
	}

	return dataCompany, nil
}

// ----------------------------------------------------------------------------------
func (repo *FavQuery) CreateFavorit(input favorit.Core) (uint, error) {
	// simpan ke DB
	newFavGorm := CoreToModel(input)

	tx := repo.db.Create(&newFavGorm) // proses query insert
	if tx.Error != nil {
		return 0, tx.Error
	}

	return newFavGorm.ID, nil
}

func (repo *FavQuery) GetAllFavorit(userID uint) ([]favorit.Core, error) {
	var productsDataGorm []database.Favourite
	tx := repo.db.Where("jobseeker_id = ?", userID).Find(&productsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProductCore := ModelGormToCore(productsDataGorm)

	return allProductCore, nil
}

func (repo *FavQuery) DeleteFavById(input []favorit.Core, id int) error {
	favGorm := CoretoModelGorm(input)

	txDel := repo.db.Delete(&favGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}

// func (repo *FavQuery) DeleteFavById(JobId uint) error {
// 	var user database.Favorites
// 	tx := repo.db.DeleteFavById(&user, JobId)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("data not found")
// 	}
// 	return nil
// }
