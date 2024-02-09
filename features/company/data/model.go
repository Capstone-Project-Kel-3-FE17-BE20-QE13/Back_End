package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/company"
)

func CoreCompannyToModel(input company.CompanyCore) database.Company {
	return database.Company{
		Company_name: input.Company_name,
		Full_name:    input.Full_name,
		Email:        input.Email,
		Password:     input.Password,
		Company_type: input.Company_type,
		Company_size: input.Company_size,
		Website:      input.Website,
	}
}
