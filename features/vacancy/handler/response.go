package handler

import (
	"JobHuntz/features/vacancy"
)

type JobResponse struct {
	ID              uint   `json:"id" form:"id"`
	CompanyID       uint   `json:"company_id" form:"company_id"`
	Name            string `json:"name" form:"name"`
	Job_type        string `json:"job_type" form:"job_type"`
	Salary_range    string `json:"salary_range" form:"salary_range"`
	Category        string `json:"category" form:"category"`
	Job_description string `json:"job_desc" form:"job_desc"`
	Job_requirement string `json:"job_req" form:"job_req"`
	Created_by      uint   `json:"created_by" form:"created_by"`
	// Applications []_appResponses.ApplicationResponse `json:"applications"`

}

func FromCore(input vacancy.Core) JobResponse {
	return JobResponse{
		ID:              input.ID,
		CompanyID:       input.CompanyID,
		Name:            input.Name,
		Job_type:        input.Job_type,
		Salary_range:    input.Salary_range,
		Category:        input.Category,
		Job_description: input.Job_description,
		Job_requirement: input.Job_requirement,
		Created_by:      input.Created_by,
		// Applications: _appResponses.ListFromDomain(domain.Applications),

	}
}

func ListFromCore(domain []vacancy.Core) (response []JobResponse) {
	for _, job := range domain {
		response = append(response, FromCore(job))
	}
	return
}
