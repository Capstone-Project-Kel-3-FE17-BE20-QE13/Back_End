package handler

import "JobHuntz/features/application"

type ApplyRequest struct {
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	JobId       uint   `json:"job_id" form:"job_id"`
	Status      string `json:"status" form:"status"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapApplyReqToCoreApply(req ApplyRequest) application.Core {
	return application.Core{
		JobseekerID: req.JobseekerID,
		JobId:       req.JobId,
		Status:      req.Status,
	}
}
