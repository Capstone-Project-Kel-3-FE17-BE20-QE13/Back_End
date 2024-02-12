package handler

import (
	"JobHuntz/features/application"
)

type ApplyResponse struct {
	ID                 uint   `json:"id" form:"id"`
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Position           string `json:"position" form:"position"`
	Company_name       string `json:"company_name" form:"company_name"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreApplyToApplyRes(input application.Core) ApplyResponse {
	return ApplyResponse{
		ID:                 input.ID,
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Position:           input.Position,
		Company_name:       input.Company_name,
		Status_application: input.Status_application,
	}
}
