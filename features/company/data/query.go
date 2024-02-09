package data

import (
	"JobHuntz/features/company"
	"JobHuntz/utils/responses"

	"gorm.io/gorm"
)

type CompanyQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) company.CompanyDataInterface {
	return &CompanyQuery{
		db: db,
	}
}

func (repo *CompanyQuery) RegisterCompany(input company.CompanyCore) error {
	newCompany := CoreCompannyToModel(input)
	newCompany.Password = responses.HashPassword(input.Password)
	newCompany.Status_Verification = "Unverified"

	tx := repo.db.Create(&newCompany)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
