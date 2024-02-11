package company

import "mime/multipart"

type CompanyCore struct {
	ID                  uint
	Company_name        string `validate:"required"`
	Full_name           string `validate:"required"`
	Email               string `validate:"required"`
	Password            string `validate:"required"`
	Company_type        string `validate:"required"`
	Company_size        string `validate:"required"`
	Website             string `validate:"required"`
	Description         string
	Status_Verification string
	Banners             string
	Address             string
	Phone               string
	Token               string
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
