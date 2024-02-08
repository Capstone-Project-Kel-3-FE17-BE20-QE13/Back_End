package data

import (
	"JobHuntz/features/user"
	"JobHuntz/utils/responses"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

func (repo *UserQuery) Register(input user.Core) error {
	// simpan ke DB
	newUserGorm := CoreUserToModel(input)
	newUserGorm.Status_Verification = "Unverified"
	newUserGorm.Password = responses.HashPassword(input.Password)

	tx := repo.db.Create(&newUserGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *UserQuery) AddCareer(input user.CareerCore) error {
	// simpan ke DB
	newCareerGorm := CoreCareerToModel(input)

	tx := repo.db.Create(&newCareerGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
