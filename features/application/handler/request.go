package handler

import "JobHuntz/features/application"

type ApplyRequest struct {
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapApplyReqToCoreApply(input ApplyRequest) application.Core {
	return application.Core{
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Status_application: input.Status_application,
	}
}
