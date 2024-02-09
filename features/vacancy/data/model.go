package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/vacancy"
)

func CoreToModel(input vacancy.Core) database.Job {
	return database.Job{
		// ID:          job.ID,
		Name:        input.Name,
		TypeJob:     input.TypeJob,
		Salary:      input.Salary,
		JobDesc:     input.JobDesc,
		Requirement: input.Requirement,
		CreatedBy:   input.CreatedBy,
		CompanyId:   input.CompanyId,
		Category:    input.Category,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),
	}
}

func CoretoModelGorm(data []vacancy.Core) []database.Job {
	var jobDataGorm []database.Job
	for _, input := range data {
		var jobGorm = database.Job{
			// ID:          job.ID,
			Name:        input.Name,
			TypeJob:     input.TypeJob,
			Salary:      input.Salary,
			JobDesc:     input.JobDesc,
			Requirement: input.Requirement,
			CreatedBy:   input.CreatedBy,
			CompanyId:   input.CompanyId,
			Category:    input.Category,
			// Company:      input.Company.ToDomain(),
			// Applications: applications.ListToCore(job.Applications),

		}
		jobDataGorm = append(jobDataGorm, jobGorm)
	}

	return jobDataGorm
}

func ModelToCore(input database.Job) vacancy.Core {
	return vacancy.Core{
		ID:          input.ID,
		Name:        input.Name,
		TypeJob:     input.TypeJob,
		Salary:      input.Salary,
		JobDesc:     input.JobDesc,
		Requirement: input.Requirement,
		CreatedBy:   input.CreatedBy,
		CompanyId:   input.CompanyId,
		Category:    input.Category,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),

	}
}

func ModelGormToCore(data []database.Job) []vacancy.Core {
	var jobData []vacancy.Core
	for _, input := range data {
		var jobInput = vacancy.Core{
			ID:          input.ID,
			Name:        input.Name,
			TypeJob:     input.TypeJob,
			Salary:      input.Salary,
			JobDesc:     input.JobDesc,
			Requirement: input.Requirement,
			CreatedBy:   input.CreatedBy,
			CompanyId:   input.CompanyId,
			Category:    input.Category,
			// Company:      input.Company.ToDomain(),
			// Applications: applications.ListToCore(job.Applications),

		}
		jobData = append(jobData, jobInput)
	}

	return jobData
}
