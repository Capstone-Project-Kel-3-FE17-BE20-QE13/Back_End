package data

import (
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"

	"gorm.io/gorm"
)

type JobseekerQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) jobseeker.JobseekerDataInterface {
	return &JobseekerQuery{
		db: db,
	}
}

func (repo *JobseekerQuery) Register(input jobseeker.JobseekerCore) error {
	newSeekerGorm := CoreJobseekerToModel(input)
	newSeekerGorm.Password = responses.HashPassword(input.Password)
	newSeekerGorm.Status_Verification = "Unverified"

	tx := repo.db.Create(&newSeekerGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// // Login implements user.UserDataInterface.
// func (repo *UserQuery) Login(email string) (user.UserCore, error) {
// 	var dataUser database.User
// 	tx := repo.db.Where("email = ?", email).First(&dataUser)
// 	if tx.Error != nil {
// 		return user.UserCore{}, errors.New(tx.Error.Error() + ", invalid email")
// 	}

// 	if tx.RowsAffected == 0 {
// 		return user.UserCore{}, errors.New("login failed, invalid email")
// 	}

// 	userCore := ModelUserToCore(dataUser)
// 	return userCore, nil
// }

// func (repo *UserQuery) AddCareer(input user.CareerCore) error {
// 	// simpan ke DB
// 	newCareerGorm := CoreCareerToModel(input)

// 	tx := repo.db.Create(&newCareerGorm) // proses query insert
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }
