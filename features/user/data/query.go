package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/user"
	"JobHuntz/utils/responses"
	"log"

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

func (repo *UserQuery) AddVerif(status string, email string) error {
	var newSeeker database.Jobseeker
	newSeeker.Status_Verification = status

	if err := repo.db.Model(&database.User{}).Where("email = ?", email).Select("id").Scan(&newSeeker.UserID).Error; err != nil {
		log.Println("Error:", err)
		return err
	}

	tx := repo.db.Create(&newSeeker) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *UserQuery) Register(input user.UserCore) error {
	newUserGorm := CoreUserToModel(input)
	newUserGorm.Password = responses.HashPassword(input.Password)

	tx := repo.db.Create(&newUserGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	var newSeekerStat database.Jobseeker
	newSeekerStat.Status_Verification = "Unverified"
	repo.AddVerif(newSeekerStat.Status_Verification, newUserGorm.Email)

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
