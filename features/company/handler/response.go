package handler

import (
	"JobHuntz/features/company"
)

type CompanyRespon struct {
	ID                  uint   `json:"id"`
	Company_name        string `json:"company_name"`
	Full_name           string `json:"full_name"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	Address             string `json:"address"`
	Phone               string `json:"phone"`
	Company_type        string `json:"company_type"`
	Company_size        string `json:"company_size"`
	Website             string `json:"website"`
	Description         string `json:"description"`
	Status_Verification string `json:"status_verification"`
	Banners             string `json:"banners"`
	Token               string `json:"token"`
}

type CompanyResponById struct {
	ID                  uint   `json:"id"`
	Company_name        string `json:"company_name"`
	Full_name           string `json:"full_name"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	Phone               string `json:"phone"`
	Company_type        string `json:"company_type"`
	Company_size        string `json:"company_size"`
	Website             string `json:"website"`
	Description         string `json:"description"`
	Status_Verification string `json:"status_verification"`
	Banners             string `json:"banners"`
}

type LoginResponse struct {
	Email string `gorm:"not null;unique" json:"email" form:"email"`
	Roles string `gorm:"not null;unique" json:"roles" form:"roles"`
	Token string `json:"token" form:"token"`
}

func ResponCompanyToResponseLogin(input company.CompanyCore, token string) LoginResponse {
	return LoginResponse{
		Email: input.Email,
		Roles: input.Roles,
		Token: token,
	}
}

func CoreResponGetByid(data company.CompanyCore) CompanyResponById {
	return CompanyResponById{
		ID:                  data.ID,
		Company_name:        data.Company_name,
		Full_name:           data.Full_name,
		Email:               data.Email,
		Address:             data.Address,
		Phone:               data.Phone,
		Company_type:        data.Company_type,
		Company_size:        data.Company_size,
		Website:             data.Website,
		Description:         data.Description,
		Status_Verification: data.Status_Verification,
		Banners:             data.Banners,
	}
}
