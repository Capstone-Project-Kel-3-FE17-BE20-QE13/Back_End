package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
	"JobHuntz/features/verification"

	"gorm.io/gorm"
)

type VerificationQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) verification.VerificationDataInterface {
	return &VerificationQuery{
		db: db,
	}
}

func (repo *VerificationQuery) GetDataJobseeker(userID uint) (jobseeker.JobseekerCore, error) {
	var data database.Jobseeker

	tx := repo.db.Where("id = ?", userID).First(&data)
	if tx.Error != nil {
		return jobseeker.JobseekerCore{}, tx.Error
	}

	dataCore := ModJobseekerToCore(data)

	return dataCore, nil
}

func (repo *VerificationQuery) GetDataCompany(userID uint) (company.CompanyCore, error) {
	var data database.Company

	tx := repo.db.Where("id = ?", userID).First(&data)
	if tx.Error != nil {
		return company.CompanyCore{}, tx.Error
	}

	dataCore := ModeCompanyToCore(data)

	return dataCore, nil
}

func (repo *VerificationQuery) AddOrder(input verification.OrderCore) error {
	newOrder := CoreOrderToModel(input)

	tx := repo.db.Create(&newOrder) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// func (repo *VerificationQuery) AddOrderCompany(input verification.OrderCompanyCore) error {
// 	newOrder := CoreCompanyToModel(input)

// 	tx := repo.db.Create(&newOrder) // proses query insert
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }
