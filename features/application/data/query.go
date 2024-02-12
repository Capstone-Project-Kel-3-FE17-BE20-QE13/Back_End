package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"
	"JobHuntz/features/favorit"
	"database/sql"

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

func (repo *ApplyQuery) GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error) {
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

func (repo *ApplyQuery) CreateApplication(input application.Core) (uint, error) {
	// simpan ke DB
	newApplyGorm := CoreToModel(input)

	tx := repo.db.Create(&newApplyGorm) // proses query insert
	if tx.Error != nil {
		return 0, tx.Error
	}

	return newApplyGorm.ID, nil
}

func (repo *ApplyQuery) GetAllApplications(jobseekerID uint) ([]application.Core, error) {
	var applicationsDataGorm []database.Application
	tx := repo.db.Where("jobseeker_id = ?", jobseekerID).Find(&applicationsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allApplicationsCore := ModelGormToCore(applicationsDataGorm)

	return allApplicationsCore, nil
}
