package handler

import (
	"JobHuntz/features/jobs"
	"time"
)

type JobResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"Name"`
	TypeJob string
	Salary  string `json:"Salary"`
	// CategoryId   int                                 `json:"categoryId"`
	// Category     _catResponses.CategoryResponse      `json:"category"`
	JobDesc     string `json:"jobDesc"`
	Requirement string `json:"requirement"`
	CreatedBy   uint   `json:"createdBy"`
	CompanyId   uint   `json:"companyId"`
	// Applications []_appResponses.ApplicationResponse `json:"applications"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromCore(domain jobs.Core) JobResponse {
	return JobResponse{
		ID:   domain.ID,
		Name: domain.Name,
		// CategoryId:   domain.CategoryId,
		// Category:     _catResponses.FromDomain(domain.Category),
		JobDesc:     domain.JobDesc,
		Requirement: domain.Requirement,
		CreatedBy:   domain.CreatedBy,
		CompanyId:   domain.CompanyId,
		// Applications: _appResponses.ListFromDomain(domain.Applications),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromCore(domain []jobs.Core) (response []JobResponse) {
	for _, job := range domain {
		response = append(response, FromCore(job))
	}
	return
}
