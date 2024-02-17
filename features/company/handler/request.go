package handler

import "JobHuntz/features/company"

type CompanyRequest struct {
	Company_name string
	Full_name    string
	Email        string
	Password     string
	Company_type string
	Company_size string
	Website      string
}

type CompanyRequestUpdate struct {
	Company_name string `json:"company_name" form:"company_name"`
	Email        string `json:"email" form:"email"`
	Full_name    string `json:"full_name" form:"full_name"`
	Address      string `json:"address" form:"address"`
	Phone        string `json:"phone" form:"phone" validate:"numeric"`
	Company_size string `json:"company_size" form:"company_size"`
	Website      string `json:"website" form:"website"`
	Description  string `json:"description" form:"description"`
	Banners      string `json:"banners" form:"banners"`
}

func RequestCompanyUpdateToCore(input CompanyRequestUpdate) company.CompanyCore {
	return company.CompanyCore{
		Company_name: input.Company_name,
		Full_name:    input.Full_name,
		Email:        input.Email,
		Company_size: input.Company_size,
		Website:      input.Website,
		Description:  input.Description,
		Banners:      input.Banners,
		Address:      input.Address,
		Phone:        input.Phone,
	}
}

type CompanyRequestLogin struct {
	Email    string
	Password string
}

func RequestCompanyToCore(input CompanyRequest) company.CompanyCore {
	return company.CompanyCore{
		Company_name: input.Company_name,
		Full_name:    input.Full_name,
		Email:        input.Email,
		Password:     input.Password,
		Company_type: input.Company_type,
		Company_size: input.Company_size,
		Website:      input.Website,
	}
}
