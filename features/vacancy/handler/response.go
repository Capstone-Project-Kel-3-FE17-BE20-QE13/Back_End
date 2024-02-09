package handler

import (
	"JobHuntz/features/vacancy"
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

}

func FromCore(domain vacancy.Core) JobResponse {
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

	}
}

func ListFromCore(domain []vacancy.Core) (response []JobResponse) {
	for _, job := range domain {
		response = append(response, FromCore(job))
	}
	return
}
