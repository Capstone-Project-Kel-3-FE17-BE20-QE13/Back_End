package handler

import (
	"JobHuntz/features/vacancy"
)

type JobResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"Name"`
	TypeJob     string `json:"type"`
	Salary      string `json:"Salary"`
	Category    string `json:"category"`
	JobDesc     string `json:"jobDesc"`
	Requirement string `json:"requirement"`
	CompanyId   uint   `json:"companyId"`
	// Applications []_appResponses.ApplicationResponse `json:"applications"`

}

func FromCore(domain vacancy.Core) JobResponse {
	return JobResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		TypeJob:     domain.TypeJob,
		Salary:      domain.Salary,
		Category:    domain.Category,
		JobDesc:     domain.JobDesc,
		Requirement: domain.Requirement,
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
