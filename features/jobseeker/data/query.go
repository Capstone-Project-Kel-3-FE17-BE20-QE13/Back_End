package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
	newSeekerGorm.Birth_date = time.Date(1700, 1, 1, 0, 0, 0, 0, time.UTC)
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

func (repo *JobseekerQuery) UpdateProfile(seekerID uint, data jobseeker.JobseekerCore) error {
	newUpdateGorm := CoreJobseekerToModel(data)

	txUpdates := repo.db.Model(&database.Jobseeker{}).Where("id = ?", seekerID).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) CV(fileHeader *multipart.FileHeader) (*uploader.UploadResult, error) {
	urlCloudinary := "cloudinary://377166738273893:ga3Zq7Ts84gJ-Ltn-gyMkTgHd40@dltcy9ghn"

	file, errHeader := fileHeader.Open()
	if errHeader != nil {
		return nil, errors.New(errHeader.Error() + "cannot open fileHeader")
	}

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	resp, errUpload := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if errUpload != nil {
		return nil, errors.New(errUpload.Error() + "cannot upload file")
	}

	return resp, nil
}

func (repo *JobseekerQuery) AddCV(input jobseeker.CVCore) error {
	newCV := CoreCVToModel(input)

	tx := repo.db.Create(&newCV) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) ReadCV(seekerID uint) (jobseeker.CVCore, error) {
	var singleCVGorm database.CV
	tx := repo.db.First(&singleCVGorm, seekerID)
	if tx.Error != nil {
		return jobseeker.CVCore{}, errors.New(tx.Error.Error() + "cannot get data of cv")
	}

	singleCVCore := ModelCVToCore(singleCVGorm)

	return singleCVCore, nil
}

func (repo *JobseekerQuery) UpdateCV(input jobseeker.CVCore) error {
	newCVGorm := CoreCVToModel(input)

	txUpdates := repo.db.Model(&database.CV{}).Where("jobseeker_id = ?", newCVGorm.JobseekerID).Updates(newCVGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveCV(input uint) error {
	result := repo.db.Where("jobseeker_id = ?", input).Delete(&database.CV{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete cv")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}

func (repo *JobseekerQuery) AddCareer(input jobseeker.CareerCore) error {
	// simpan ke DB
	newCareerGorm := CoreCareerToModel(input)

	tx := repo.db.Create(&newCareerGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) GetCareerByID(input uint) (jobseeker.CareerCore, error) {
	var singleCareerGorm database.Career
	tx := repo.db.First(&singleCareerGorm, input)
	if tx.Error != nil {
		return jobseeker.CareerCore{}, tx.Error
	}

	singleCareerCore := ModelCareerToCore(singleCareerGorm)

	return singleCareerCore, nil
}

func (repo *JobseekerQuery) GetCareerList(input uint) ([]jobseeker.CareerCore, error) {
	var careersDataGorm []database.Career
	tx := repo.db.Where("jobseeker_id = ?", input).Find(&careersDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allCareersCore := ModelCareersToCore(careersDataGorm)

	return allCareersCore, nil
}

func (repo *JobseekerQuery) UpdateCareer(career_id uint, input jobseeker.CareerCore) error {
	newCareerGorm := CoreCareerToModel(input)

	txUpdates := repo.db.Model(&database.Career{}).Where("id = ?", career_id).Updates(newCareerGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveCareer(input uint) error {
	result := repo.db.Where("id = ?", input).Delete(&database.Career{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete career")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}
