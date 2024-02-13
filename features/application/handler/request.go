package handler

import "JobHuntz/features/application"

type ApplyRequest struct {
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Position           string `json:"position" form:"position"`
	Company_name       string `json:"company_name" form:"company_name"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

type ApplicationRequestStatus struct {
	Status_application string `json:"stat_app" form:"stat_app"`
}

func RequestToCore(input ApplicationRequestStatus) application.Core {
	return application.Core{
		Status_application: input.Status_application,
	}
}

// Mapping dari struct TaskRequest To struct Core Task
func MapApplyReqToCoreApply(input ApplyRequest) application.Core {
	return application.Core{
		JobseekerID:        input.JobseekerID,
		VacancyID:          input.VacancyID,
		Position:           input.Position,
		Company_name:       input.Company_name,
		Status_application: input.Status_application,
	}
}
