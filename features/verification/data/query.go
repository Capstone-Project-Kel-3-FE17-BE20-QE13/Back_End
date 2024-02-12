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

func (repo *VerificationQuery) AddOrder(input verification.OrderCore) error {
	newHistory := HistoryToModel(input)

	tx := repo.db.Create(&newHistory) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
