package service

import (
	"JobHuntz/features/application"
	"JobHuntz/features/favorit"
	"database/sql"
	"errors"
)

type ApplyService struct {
	Repo application.ApplyDataInterface
}

func New(repo application.ApplyDataInterface) application.ApplyServiceInterface {
	return &ApplyService{
		Repo: repo,
	}
}

func (uc *ApplyService) GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error) {

	application, err := uc.Repo.GetDataCompany(dbRaw, vacancyID)

	if err != nil {
		return favorit.DataCompanyCore{}, err
	}

	return application, nil
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

func (uc *ApplyService) GetAllApplicationsCompany(VacancyID uint) ([]application.Core, error) {
	result, err := uc.Repo.GetAllApplications(VacancyID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *ApplyService) EditApplication(id uint, input application.Core) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.Repo.Edit(id, input)
	return err
}
