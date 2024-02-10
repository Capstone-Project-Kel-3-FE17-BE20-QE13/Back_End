package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/company"
	"JobHuntz/utils/responses"
	"errors"

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

func (repo *CompanyQuery) LoginCompany(email string) (company.CompanyCore, error) {
	var dataCompany database.Company
	tx := repo.db.Where("email = ?", email).First(&dataCompany)
	if tx.Error != nil {
		return company.CompanyCore{}, errors.New(tx.Error.Error() + ", invalid email")
	}

	if tx.RowsAffected == 0 {
		return company.CompanyCore{}, errors.New("login failed, invalid email")
	}

	companyCore := CoreLoginCompanyToModel(dataCompany)
	return companyCore, nil
}
