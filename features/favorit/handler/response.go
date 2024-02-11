package handler

import "JobHuntz/features/favorit"

type FavResponse struct {
	ID           uint   `json:"id" form:"id"`
	JobseekerID  uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID    uint   `json:"vacancy_id" form:"vacancy_id"`
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreApplyToApplyRes(input favorit.Core) FavResponse {
	return FavResponse{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		VacancyID:    input.VacancyID,
		Position:     input.Position,
		Company_name: input.Company_name,
	}
}
