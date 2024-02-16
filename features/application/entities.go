package application

import (
	"JobHuntz/features/favorit"
	"JobHuntz/features/jobseeker"
	"database/sql"
)

type Core struct {
	ID                 uint   `json:"id" form:"id"`
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Position           string `json:"position" form:"position"`
	Company_name       string `json:"company_name" form:"company_name"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

type ListApplicantsCore struct {
	ID                 uint   `json:"id" form:"id"`
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Full_name          string `json:"full_name" form:"full_name"`
	Username           string `json:"username" form:"username"`
	Email              string `json:"email" form:"email"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Position           string `json:"position" form:"position"`
	Company_name       string `json:"company_name" form:"company_name"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

type ApplyDataInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error)
	GetMyData(userID uint) (jobseeker.JobseekerCore, error)
	CountApplication(dbRaw *sql.DB, userID uint) (uint, error)
	CreateApplication(input Core) error
	GetAllApplications(JobseekerID uint) ([]Core, error)
	GetAllApplicationsCompany(VacancyID uint) ([]Core, error)
	Edit(id uint, input Core) error
}

type ApplyServiceInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error)
	GetMyData(userID uint) (jobseeker.JobseekerCore, error)
	CountApplication(dbRaw *sql.DB, userID uint) (uint, error)
	CreateApplication(input Core, count uint, status string) error
	GetAllApplications(JobseekerID uint) ([]Core, error)
	GetAllApplicationsCompany(VacancyID uint) ([]Core, error)
	EditApplication(id uint, input Core) error
}
