package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/vacancy"
)

func CoreToModel(input vacancy.Core) database.Vacancy {
	return database.Vacancy{
		// ID:          job.ID,
		CompanyID:       input.CompanyID,
		Name:            input.Name,
		Job_type:        input.Job_type,
		Salary_range:    input.Salary_range,
		Category:        input.Category,
		Job_description: input.Job_description,
		Job_requirement: input.Job_requirement,
		Created_by:      input.Created_by,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),
	}
}

func CoretoModelGorm(data []vacancy.Core) []database.Vacancy {
	var jobDataGorm []database.Vacancy
	for _, input := range data {
		var jobGorm = database.Vacancy{
			// ID:          job.ID,
			CompanyID:       input.CompanyID,
			Name:            input.Name,
			Job_type:        input.Job_type,
			Salary_range:    input.Salary_range,
			Category:        input.Category,
			Job_description: input.Job_description,
			Job_requirement: input.Job_requirement,
			Created_by:      input.Created_by,
		}
		jobDataGorm = append(jobDataGorm, jobGorm)
	}

	return jobDataGorm
}

func ModelToCore(input database.Vacancy) vacancy.Core {
	return vacancy.Core{
		ID:              input.ID,
		CompanyID:       input.CompanyID,
		Name:            input.Name,
		Job_type:        input.Job_type,
		Salary_range:    input.Salary_range,
		Category:        input.Category,
		Job_description: input.Job_description,
		Job_requirement: input.Job_requirement,
		Created_by:      input.Created_by,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),

	}
}

func ModelGormToCore(data []database.Vacancy) []vacancy.Core {
	var jobData []vacancy.Core
	for _, input := range data {
		var jobInput = vacancy.Core{
			ID:              input.ID,
			CompanyID:       input.CompanyID,
			Name:            input.Name,
			Job_type:        input.Job_type,
			Salary_range:    input.Salary_range,
			Category:        input.Category,
			Job_description: input.Job_description,
			Job_requirement: input.Job_requirement,
			Created_by:      input.Created_by,
			// Company:      input.Company.ToDomain(),
			// Applications: applications.ListToCore(job.Applications),

		}
		jobData = append(jobData, jobInput)
	}

	return jobData
}
