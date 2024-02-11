package handler

import "JobHuntz/features/favorit"

type FavRequest struct {
	JobseekerID  uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID    uint   `json:"vacancy_id" form:"vacancy_id"`
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapApplyReqToCoreApply(input FavRequest) favorit.Core {
	return favorit.Core{
		JobseekerID:  input.JobseekerID,
		VacancyID:    input.VacancyID,
		Position:     input.Position,
		Company_name: input.Company_name,
	}
}
