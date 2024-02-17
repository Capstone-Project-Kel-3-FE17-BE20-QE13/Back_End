package favorit

import (
	"database/sql"
)

type Core struct {
	ID           uint   `json:"id" form:"id"`
	JobseekerID  uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID    uint   `json:"vacancy_id" form:"vacancy_id"`
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
}

type DataCompanyCore struct {
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
}

type FavDataInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (DataCompanyCore, error)
	CreateFavorit(input Core) (uint, error)
	GetAllFavorit(userID uint) ([]Core, error)
	// DeleteFavById(JobId uint) error
	DeleteFavById(input []Core, id int) error
}

type FavServiceInterface interface {
	GetDataCompany(dbRaw *sql.DB, vacancyID uint) (DataCompanyCore, error)
	CreateFavorit(input Core) (uint, error)
	GetAllFavorit(userID uint) ([]Core, error)
	// DeleteFavById(JobId uint) error
	DeleteFavById(input []Core, id int) error
}
