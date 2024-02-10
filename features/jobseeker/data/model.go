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

func CoreCVToModel(input jobseeker.CVCore) database.CV {
	return database.CV{
		JobseekerID: input.JobseekerID,
		CV_file:     input.CV_file,
	}
}

func CoreCareerToModel(input jobseeker.CareerCore) database.Career {
	return database.Career{
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}

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

func ModelCVToCore(input database.CV) jobseeker.CVCore {
	return jobseeker.CVCore{
		JobseekerID: input.JobseekerID,
		CV_file:     input.CV_file,
	}
}

func ModelCareerToCore(input database.Career) jobseeker.CareerCore {
	return jobseeker.CareerCore{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}

func ModelCareersToCore(data []database.Career) []jobseeker.CareerCore {
	var careersData []jobseeker.CareerCore
	for _, input := range data {
		var careerInput = jobseeker.CareerCore{
			ID:           input.ID,
			JobseekerID:  input.JobseekerID,
			Position:     input.Position,
			Company_name: input.Company_name,
			Date_start:   input.Date_start,
			Date_end:     input.Date_end,
		}
		careersData = append(careersData, careerInput)
	}

	return careersData
}
