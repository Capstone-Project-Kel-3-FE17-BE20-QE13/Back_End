package data

import (
	"JobHuntz/features/category"
	"JobHuntz/features/jobs"
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Name        string
	TypeJob     string
	Salary      string
	CategoryId  int
	Category    category.Core
	JobDesc     string
	Requirement string
	CreatedBy   uint
	CompanyId   uint
	// Company      companies.Domain
	// Applications []applications.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToModel(input jobs.Core) Job {
	return Job{
		// ID:          job.ID,
		Name:        input.Name,
		TypeJob:     input.TypeJob,
		Salary:      input.Salary,
		JobDesc:     input.JobDesc,
		Requirement: input.Requirement,
		CreatedBy:   input.CreatedBy,
		CompanyId:   input.CompanyId,
		CategoryId:  input.CategoryId,
		Category:    input.Category,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func CoretoModelGorm(data []jobs.Core) []Job {
	var jobDataGorm []Job
	for _, input := range data {
		var jobGorm = Job{
			// ID:          job.ID,
			Name:        input.Name,
			TypeJob:     input.TypeJob,
			Salary:      input.Salary,
			JobDesc:     input.JobDesc,
			Requirement: input.Requirement,
			CreatedBy:   input.CreatedBy,
			CompanyId:   input.CompanyId,
			CategoryId:  input.CategoryId,
			Category:    input.Category,
			// Company:      input.Company.ToDomain(),
			// Applications: applications.ListToCore(job.Applications),
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		}
		jobDataGorm = append(jobDataGorm, jobGorm)
	}

	return jobDataGorm
}

func ModelToCore(input Job) jobs.Core {
	return jobs.Core{
		ID:          input.ID,
		Name:        input.Name,
		TypeJob:     input.TypeJob,
		Salary:      input.Salary,
		JobDesc:     input.JobDesc,
		Requirement: input.Requirement,
		CreatedBy:   input.CreatedBy,
		CompanyId:   input.CompanyId,
		CategoryId:  input.CategoryId,
		Category:    input.Category,
		// Company:      input.Company.ToDomain(),
		// Applications: applications.ListToCore(job.Applications),
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ModelGormToCore(data []Job) []jobs.Core {
	var jobData []jobs.Core
	for _, input := range data {
		var jobInput = jobs.Core{
			ID:          input.ID,
			Name:        input.Name,
			TypeJob:     input.TypeJob,
			Salary:      input.Salary,
			JobDesc:     input.JobDesc,
			Requirement: input.Requirement,
			CreatedBy:   input.CreatedBy,
			CompanyId:   input.CompanyId,
			CategoryId:  input.CategoryId,
			Category:    input.Category,
			// Company:      input.Company.ToDomain(),
			// Applications: applications.ListToCore(job.Applications),
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		}
		jobData = append(jobData, jobInput)
	}

	return jobData
}
