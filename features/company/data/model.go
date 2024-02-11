package data

import (
	"JobHuntz/features/company"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID                  uint
	Company_name        string `gorm:"not null" json:"company_name" form:"company_name"`
	Full_name           string `json:"full_name" form:"full_name"`
	Email               string `gorm:"not null" json:"email" form:"email"`
	Password            string `gorm:"not null" json:"password" form:"password"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Company_type        string `gorm:"not null" json:"company_type" form:"company_type"`
	Company_size        string `gorm:"not null" json:"company_size" form:"company_size"`
	Website             string `gorm:"not null" json:"website" form:"website"`
	Description         string `json:"description" form:"description"`
	Status_Verification string `json:"status_verification" form:"status_verification"`
	Banners             string `json:"banners" form:"banners"`
}

func (u Company) ModelRegisterToCore() company.CompanyCore {
	return company.CompanyCore{
		ID:           u.ID,
		Company_name: u.Company_name,
		Full_name:    u.Full_name,
		Email:        u.Email,
		Password:     u.Password,
		Company_type: u.Company_type,
		Company_size: u.Company_size,
		Website:      u.Website,
	}
}

func CoreModelCompanyUpdate(input company.CompanyCore) Company {
	return Company{
		Company_name: input.Company_name,
		Full_name:    input.Full_name,
		Email:        input.Email,
		Address:      input.Address,
		Phone:        input.Phone,
		Company_type: input.Company_type,
		Company_size: input.Company_size,
		Website:      input.Website,
		Description:  input.Description,
		Banners:      input.Banners,
	}
}
