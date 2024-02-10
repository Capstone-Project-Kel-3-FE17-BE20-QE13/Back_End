package handler

import (
	"JobHuntz/features/application"
)

type ApplyResponse struct {
	JobseekerID uint
	JobId       uint
	Status      string
}

// Mapping CorePrject to TaskResponsee
func MapCoreApplyToApplyRes(core application.Core) ApplyResponse {
	return ApplyResponse{
		JobseekerID: core.JobseekerID,
		JobId:       core.JobId,
		Status:      core.Status,
	}
}
