package handler

import "JobHuntz/features/favorit"

type FavRequest struct {
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	JobId       uint   `json:"job_id" form:"job_id"`
	Name        string `json:"name" form:"name"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapApplyReqToCoreApply(req FavRequest) favorit.Core {
	return favorit.Core{
		JobseekerID: req.JobseekerID,
		JobId:       req.JobId,
		Name:        req.Name,
	}
}
