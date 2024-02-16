package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
	"JobHuntz/features/verification"
)

func ModJobseekerToCore(input database.Jobseeker) jobseeker.JobseekerCore {
	return jobseeker.JobseekerCore{
		ID:                  input.ID,
		Full_name:           input.Full_name,
		Username:            input.Username,
		Email:               input.Email,
		Password:            input.Password,
		Address:             input.Address,
		Phone:               input.Phone,
		Birth_date:          input.Birth_date,
		Gender:              input.Gender,
		Resume:              input.Resume,
		Status_Verification: input.Status_Verification,
	}
}

func ModeCompanyToCore(input database.Company) company.CompanyCore {
	return company.CompanyCore{
		ID:           input.ID,
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

func CoreOrderJobseekerToModel(input verification.OrderJobseekerCore) database.OrderJobseeker {
	return database.OrderJobseeker{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}

func CoreOrderCompanyToModel(input verification.OrderCompanyCore) database.OrderCompany {
	return database.OrderCompany{
		ID:           input.ID,
		CompanyID:    input.CompanyID,
		Price:        input.Price,
		Status_order: input.Status_order,
	}
}
