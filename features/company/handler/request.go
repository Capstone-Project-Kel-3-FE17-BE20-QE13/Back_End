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
