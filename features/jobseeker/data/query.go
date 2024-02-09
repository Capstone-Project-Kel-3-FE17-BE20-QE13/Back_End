package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"errors"

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

// Login implements user.UserDataInterface.
func (repo *JobseekerQuery) Login(email string) (jobseeker.JobseekerCore, error) {
	var dataSeeker database.Jobseeker
	tx := repo.db.Where("email = ?", email).First(&dataSeeker)
	if tx.Error != nil {
		return jobseeker.JobseekerCore{}, errors.New(tx.Error.Error() + ", invalid email")
	}

	if tx.RowsAffected == 0 {
		return jobseeker.JobseekerCore{}, errors.New("login failed, invalid email")
	}

	userCore := ModelJobseekerToCore(dataSeeker)
	return userCore, nil
}

// func (repo *UserQuery) AddCareer(input user.CareerCore) error {
// 	// simpan ke DB
// 	newCareerGorm := CoreCareerToModel(input)

// 	tx := repo.db.Create(&newCareerGorm) // proses query insert
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }
