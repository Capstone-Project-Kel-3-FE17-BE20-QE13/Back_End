package handler

import "JobHuntz/features/jobs"

type JobRequest struct {
	Name string `json:"name"`
	// CategoryId uint   `json:"categoryId"`
	CompanyId uint   `json:"companyId"`
	JobDesc   string `json:"jobDesc"`
	CreatedBy uint   `json:"createdBy"`
}

func RequestToCore(input *JobRequest) jobs.Core {
	return jobs.Core{
		Name: input.Name,
		// CategoryId: input.CategoryId,
		CompanyId: input.CompanyId,
		JobDesc:   input.JobDesc,
		CreatedBy: input.CreatedBy,
	}
}
