package service

import (
	"JobHuntz/features/application"
)

type ApplyService struct {
	Repo application.ApplyDataInterface
}

func New(repo application.ApplyDataInterface) application.ApplyServiceInterface {
	return &ApplyService{
		Repo: repo,
	}
}
func (uc *ApplyService) CreateApplication(input application.Core) (uint, error) {

	application, err := uc.Repo.CreateApplication(input)

	if err != nil {
		return 0, err
	}

	return application, nil
}

func (uc *ApplyService) GetAllApplications(JobseekerID uint) ([]application.Core, error) {
	result, err := uc.Repo.GetAllApplications(JobseekerID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
