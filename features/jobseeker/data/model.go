package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/jobseeker"
)

func CoreJobseekerToModel(input jobseeker.JobseekerCore) database.Jobseeker {
	return database.Jobseeker{
		Full_name:           input.Full_name,
		Username:            input.Username,
		Email:               input.Email,
		Password:            input.Password,
		Address:             input.Address,
		Phone:               input.Phone,
		Birth_date:          input.Birth_date,
		Gender:              input.Gender,
		Resume:              input.Resume,
		Status_Verification: input.Status_Verification,
	}
}

// func CoreCareerToModel(input user.CareerCore) database.Career {
// 	return database.Career{
// 		JobseekerID:  input.JobseekerID,
// 		Position:     input.Position,
// 		Company_name: input.Company_name,
// 		Date_start:   input.Date_start,
// 		Date_end:     input.Date_end,
// 	}
// }

func ModelJobseekerToCore(input database.Jobseeker) jobseeker.JobseekerCore {
	return jobseeker.JobseekerCore{
		ID:                  input.ID,
		Full_name:           input.Full_name,
		Username:            input.Username,
		Email:               input.Email,
		Password:            input.Password,
		Address:             input.Address,
		Phone:               input.Phone,
		Birth_date:          input.Birth_date,
		Gender:              input.Gender,
		Resume:              input.Resume,
		Status_Verification: input.Status_Verification,
	}
}
