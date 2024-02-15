package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"
	"JobHuntz/features/jobseeker"
)

func CoreToModel(input application.Core) database.Application {
	return database.Application{
		// ID:          input.Id,
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Position:           input.Position,
		Company_name:       input.Company_name,
		Status_application: input.Status_application,
	}
}

func CoretoModelGorm(data []application.Core) []database.Application {
	var applicationsDataGorm []database.Application
	for _, input := range data {
		var applicationGorm = database.Application{
			JobseekerID:        input.JobseekerID,
			VacancyID:          input.VacancyID,
			Position:           input.Position,
			Company_name:       input.Company_name,
			Status_application: input.Status_application,
		}
		applicationsDataGorm = append(applicationsDataGorm, applicationGorm)
	}

	return applicationsDataGorm
}

func ModelToCore(input database.Application) application.Core {
	return application.Core{
		ID:                 input.ID,
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Position:           input.Position,
		Company_name:       input.Company_name,
		Status_application: input.Status_application,
	}
}

func ModelGormToCore(data []database.Application) []application.Core {
	var applicationsData []application.Core
	for _, input := range data {
		var applicationInput = application.Core{
			ID:                 input.ID,
			JobseekerID:        input.JobseekerID,
			VacancyID:          input.VacancyID,
			Position:           input.Position,
			Company_name:       input.Company_name,
			Status_application: input.Status_application,
		}
		applicationsData = append(applicationsData, applicationInput)
	}

	return applicationsData
}

func ModelGormJobseekerToCore(input database.Jobseeker) jobseeker.JobseekerCore {
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
