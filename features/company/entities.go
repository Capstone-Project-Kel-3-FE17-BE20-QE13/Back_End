package company

import (
	"mime/multipart"
)

type CompanyCore struct {
	ID                  uint   `json:"id" form:"id"`
	Company_name        string `validate:"required" gorm:"not null" json:"company_name" form:"company_name"`
	Full_name           string `validate:"required" json:"full_name" form:"full_name"`
	Email               string `validate:"required" gorm:"not null;unique" json:"email" form:"email"`
	Password            string `validate:"required" gorm:"not null" json:"password" form:"password"`
	Company_type        string `validate:"required" gorm:"not null" json:"company_type" form:"company_type"`
	Company_size        string `validate:"required" gorm:"not null" json:"company_size" form:"company_size"`
	Website             string `validate:"required" gorm:"not null" json:"website" form:"website"`
	Description         string `json:"description" form:"description"`
	Status_Verification string `json:"status_verification" form:"status_verification"`
	Banners             string `json:"banners" form:"banners"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Token               string
	Roles               string
}

type CompanyDataInterface interface {
	RegisterCompany(input CompanyCore) (data *CompanyCore, token string, err error)
	LoginCompany(email, password string) (data *CompanyCore, err error)
	GetById(id uint) (*CompanyCore, error)
	UpdateCompany(id int, input CompanyCore, file multipart.File, nameFile string) error
}

type CompanyServiceInterface interface {
	RegisterCompany(input CompanyCore) (data *CompanyCore, token string, err error)
	LoginCompany(email, password string) (data *CompanyCore, token string, err error)
	GetById(id uint) (*CompanyCore, error)
	UpdateCompany(id int, input CompanyCore, file multipart.File, nameFile string) error
}
