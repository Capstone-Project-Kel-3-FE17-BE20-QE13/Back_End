package data

import (
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

func (repo *VerificationQuery) AddOrderJobseeker(input verification.OrderJobseekerCore) error {
	newOrder := CoreJobseekerToModel(input)

	tx := repo.db.Create(&newOrder) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *VerificationQuery) AddOrderCompany(input verification.OrderCompanyCore) error {
	newOrder := CoreCompanyToModel(input)

	tx := repo.db.Create(&newOrder) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
