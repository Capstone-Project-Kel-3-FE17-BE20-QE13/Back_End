package service

import (
	"JobHuntz/features/jobs"
	"time"
)

type JobService struct {
	Repo jobs.JobDataInterface
}

func NewJob(repo jobs.JobDataInterface, timeout time.Duration) jobs.JobServiceInterface {
	return &JobService{
		Repo: repo,
	}
}
func (service *JobService) GetCompanyID(input uint) (uint, error) {
	// logic validation
	res, err := service.Repo.GetCompanyID(input)
	return res, err
}
func (service *JobService) GetJobById(input int) (jobs.Core, error) {
	// logic validation
	res, err := service.Repo.GetJobById(input)
	return res, err
}

// func (service *JobService) Photo(fileHeader *multipart.FileHeader) *uploader.UploadResult {
// 	res := service.productData.Photo(fileHeader)
// 	return res
// }

func (service *JobService) CreateJob(input jobs.Core) error {
	// logic validation
	err := service.Repo.CreateJob(input)
	return err
}

// func (service *JobService) Update(id int, input product.Core) error {
// 	//validasi
// 	if id <= 0 {
// 		return errors.New("invalid id")
// 	}
// 	//validasi inputan
// 	// ...
// 	err := service.productData.Update(id, input)
// 	return err
// }

func (service *JobService) GetAllJobs() ([]jobs.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.Repo.GetAllJobs()
	return results, err
}

func (service *JobService) DeleteJobById(input []jobs.Core, ID int) error {
	err := service.Repo.DeleteJobById(input, ID)
	return err
}

// func (service *JobService) GetSingleJob(jobID int) (jobs.Core, error) {
// 	// logic
// 	// memanggil func yg ada di data layer
// 	results, err := service.Repo.GetSingleJob(jobID)
// 	return results, err
// }

// func (service *JobService) GetStoreProducts(store_id uint) ([]product.Core, error) {
// 	// logic
// 	// memanggil func yg ada di data layer
// 	results, err := service.productData.GetStoreProducts(store_id)
// 	return results, err
// }

func (uc *JobService) SearchJobs(Name string) ([]jobs.Core, error) {
	// ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	// defer cancel()

	job, err := uc.Repo.SearchJobs(Name)
	if err != nil {
		return []jobs.Core{}, err
	}

	return job, nil
}

// func (uc *JobHandler) FilterJobByCategory(c context.Context, categoryId int) ([]jobs.Core, error) {
// 	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
// 	defer cancel()

// 	job, err := uc.Repo.FilterJobByCategory(ctx, categoryId)
// 	if err != nil {
// 		return []jobs.Core{}, err
// 	}

// 	return job, nil
// }
