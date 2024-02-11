package handler

import "JobHuntz/features/favorit"

type FavResponse struct {
	JobseekerID uint
	JobId       uint
	Name        string
}

// Mapping CorePrject to TaskResponsee
func MapCoreApplyToApplyRes(core favorit.Core) FavResponse {
	return FavResponse{
		JobseekerID: core.JobseekerID,
		JobId:       core.JobId,
		Name:        core.Name,
	}
}
