package data

import (
	// "JobHuntz/features/jobs"
	// "context"
	// "errors"

	"JobHuntz/features/jobs"
	"errors"

	"gorm.io/gorm"
)

type jobQuery struct {
	db *gorm.DB
	//contextTimeout time.Duration
}

// func (repo *jobQuery) GetStoreID(userID uint) (uint, error) {
// 	var storeData database.Store
// 	tx := repo.db.Where("user_id = ?", userID).First(&storeData)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}

// 	storeID := storeData.ID
// 	return storeID, nil
// }

// func (repo *jobQuery) Photo(fileHeader *multipart.FileHeader) *uploader.UploadResult {
// 	urlCloudinary := "cloudinary://377166738273893:ga3Zq7Ts84gJ-Ltn-gyMkTgHd40@dltcy9ghn"

// 	file, _ := fileHeader.Open()
// 	//log.Println(fileHeader.Filename)

// 	ctx := context.Background()
// 	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
// 	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
// 	//log.Println(resp.SecureURL)

// 	return resp
// }

func (repo *jobQuery) Insert(input jobs.Core) error {
	// simpan ke DB
	newProductGorm := CoreToModel(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// func (repo *jobQuery) Update(idParam int, newUpdate jobs.Core) error {
// 	newUpdateGorm := CoreToModel(newUpdate)

// 	txUpdates := repo.db.Model(&database.Product{}).Where("id = ?", idParam).Updates(newUpdateGorm)
// 	if txUpdates.Error != nil {
// 		return txUpdates.Error
// 	}

// 	return nil
// }

func (repo *jobQuery) Delete(input []jobs.Core, id int) error {
	allProductGorm := CoretoModelGorm(input)

	txDel := repo.db.Delete(&allProductGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}

func (repo *jobQuery) SelectAll() ([]jobs.Core, error) {
	var jobDataGorm []Job
	tx := repo.db.Find(&jobDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProductCore := ModelGormToCore(jobDataGorm)

	return allProductCore, nil
}

func (repo *jobQuery) GetSingleJob(id int) (jobs.Core, error) {
	var singleJobGorm Job
	tx := repo.db.First(&singleJobGorm, id)
	if tx.Error != nil {
		return jobs.Core{}, tx.Error
	}

	singleProductCore := ModelToCore(singleJobGorm)

	return singleProductCore, nil
}

// func (repo *jobQuery) GetStoreProducts(store_id uint) ([]jobs.Core, error) {
// 	var jobDataGorm []Job
// 	tx := repo.db.Where("job_id = ?", store_id).Find(&jobDataGorm) // select * from users;
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	//mapping
// 	allProductCore := ModelGormToCore(jobDataGorm)

// 	return allProductCore, nil
// }

// func New(db *gorm.DB) jobs.JobDataInterface {
// 	return &jobQuery{
// 		db: db,
// 	}
// }

// func (repo *jobQuery) GetJobById(jobID int) (jobs.Core, error) {
// 	var jobsbyid Job
// 	tx := repo.db.First(&jobsbyid, jobID)
// 	if tx.Error != nil {
// 		return jobs.Core{}, tx.Error
// 	}

// 	singleProductCore := ModelToCore(jobsbyid)

// 	return singleProductCore, nil
// }

// func (repo *jobQuery) Insert(input jobs.Core) error {
// 	// simpan ke DB
// 	newJob := CoreToModel(input)

// 	tx := repo.db.Create(&newJob) // proses query insert
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }

// func (repo *jobQuery) Update(idParam int, newUpdate jobs.Core) error {
// 	newUpdateGorm := CoreToModel(newUpdate)

// 	txUpdates := repo.db.Model(&database.Product{}).Where("id = ?", idParam).Updates(newUpdateGorm)
// 	if txUpdates.Error != nil {
// 		return txUpdates.Error
// 	}

// 	return nil
// }

// func (repo *jobQuery) Delete(input []jobs.Core, id int) error {
// 	allJob := CoretoModelGorm(input)

// 	txDel := repo.db.Delete(&allJob, id)
// 	if txDel.Error != nil {
// 		return txDel.Error
// 	}

// 	if txDel.RowsAffected == 0 {
// 		return errors.New("user's not found")
// 	}

// 	return nil
// }

// func (repo *jobQuery) SelectAll() ([]jobs.Core, error) {
// 	var jobDataGorm []database.Product
// 	tx := repo.db.Find(&jobDataGorm) // select * from users;
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	//mapping
// 	allProductCore := ModelGormToCore(jobDataGorm)

// 	return allProductCore, nil
// }
