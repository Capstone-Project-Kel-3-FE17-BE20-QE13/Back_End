package service

import (
	"JobHuntz/features/application"
	"JobHuntz/features/favorit"
	"JobHuntz/features/jobseeker"
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

func (uc *ApplyService) GetMyData(userID uint) (jobseeker.JobseekerCore, error) {
	res, err := uc.Repo.GetMyData(userID)
	if err != nil {
		return jobseeker.JobseekerCore{}, err
	}

	return res, nil
}

func (uc *ApplyService) CountApplication(dbRaw *sql.DB, userID uint) (uint, error) {
	res, err := uc.Repo.CountApplication(dbRaw, userID)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (uc *ApplyService) CreateApplication(input application.Core, count uint, status string) error {
	if status == "Unverified" && count == 3 {
		return errors.New("unverified user can't make more than 3 applications")
	} else {
		err := uc.Repo.CreateApplication(input)

		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *ApplyService) GetAllApplications(JobseekerID uint) ([]application.Core, error) {
	result, err := uc.Repo.GetAllApplications(JobseekerID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ApplyService) GetAllApplicationsCompany(VacancyID uint) ([]application.Core, error) {
	result, err := uc.Repo.GetAllApplicationsCompany(VacancyID)
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
