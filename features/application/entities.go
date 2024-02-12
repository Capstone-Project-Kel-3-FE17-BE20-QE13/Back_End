package application

import (
	"JobHuntz/features/favorit"
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

type ApplyDataInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error)
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
	GetAllApplicationsCompany(vacancyID_int int) (Core, error)
}

type ApplyServiceInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error)
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
	GetAllApplicationsCompany(vacancyID_int int) (Core, error)
}
