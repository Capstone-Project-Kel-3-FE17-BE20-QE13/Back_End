package service

import (
	"JobHuntz/features/vacancy"
)

type JobService struct {
	Repo vacancy.JobDataInterface
}

func NewJob(repo vacancy.JobDataInterface) vacancy.JobServiceInterface {
	return &JobService{
		Repo: repo,
	}
}

func (service *JobService) CountJobsByUserID(userID uint) (int, error) {
	count, err := service.Repo.CountJobsByUserID(userID)
	return count, err
}

func (service *JobService) GetById(id uint) (*vacancy.CompanyCore, error) {
	result, err := service.Repo.GetById(id)
	return result, err
}

func (service *JobService) GetJobById(input int) (vacancy.Core, error) {
	res, err := service.Repo.GetJobById(input)
	return res, err
}

func (service *JobService) CreateJob(input vacancy.Core) error {
	err := service.Repo.CreateJob(input)
	return err
}

func (service *JobService) MyCompanyVacancies(companyID uint) ([]vacancy.Core, error) {
	results, err := service.Repo.MyCompanyVacancies(companyID)
	return results, err
}

func (service *JobService) GetAllJobs() ([]vacancy.Core, error) {
	results, err := service.Repo.GetAllJobs()
	return results, err
}

func (service *JobService) DeleteJobById(input []vacancy.Core, ID int) error {
	err := service.Repo.DeleteJobById(input, ID)
	return err
}

func (service *JobService) UpdateStatus(input vacancy.JobStatusCore, vacancyID uint) error {
	err := service.Repo.UpdateStatus(input, vacancyID)
	return err
}
