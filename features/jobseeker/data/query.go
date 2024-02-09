package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"context"
	"errors"
	"mime/multipart"

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

// func (repo *UserQuery) AddCareer(input user.CareerCore) error {
// 	// simpan ke DB
// 	newCareerGorm := CoreCareerToModel(input)

// 	tx := repo.db.Create(&newCareerGorm) // proses query insert
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }