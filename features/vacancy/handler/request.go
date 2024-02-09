package handler

import "JobHuntz/features/vacancy"

type JobRequest struct {
	Name        string `json:"name"`
	TypeJob     string `json:"type_job"`
	Salary      string `json:"salary"`
	JobDesc     string `json:"jobDesc"`
	Requirement string `json:"requirement"`
	Category    string `json:"category"`
	CompanyId   uint   `json:"companyId"`
}

func RequestToCore(input *JobRequest) vacancy.Core {
	return vacancy.Core{
		Name:        input.Name,
		TypeJob:     input.TypeJob,
		Salary:      input.Salary,
		JobDesc:     input.JobDesc,
		Requirement: input.Requirement,
		Category:    input.Category,
		CompanyId:   input.CompanyId,
	}
}
