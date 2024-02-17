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
		Banners:             input.Banners,
	}
}

func CoreJobseekerRegistToModel(input jobseeker.JobseekerRegistCore) database.Jobseeker {
	return database.Jobseeker{
		Full_name: input.Full_name,
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
	}
}

func CoreJobseekerToModelUpdate(input jobseeker.JobseekerUpdateCore) database.Jobseeker {
	return database.Jobseeker{
		Full_name:  input.Full_name,
		Username:   input.Username,
		Address:    input.Address,
		Phone:      input.Phone,
		Birth_date: input.Birth_date,
		Gender:     input.Gender,
		Resume:     input.Resume,
		Banners:    input.Banners,
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

func CoreEducationToModel(input jobseeker.EducationCore) database.Education {
	return database.Education{
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}

func CoreLicenseToModel(input jobseeker.LicenseCore) database.License {
	return database.License{
		JobseekerID:    input.JobseekerID,
		License_name:   input.License_name,
		Published_date: input.Published_date,
		Expiry_date:    input.Expiry_date,
		License_file:   input.License_file,
	}
}

func CoreSkillToModel(input jobseeker.SkillCore) database.Skill {
	return database.Skill{
		JobseekerID: input.JobseekerID,
		Skill:       input.Skill,
		Description: input.Description,
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
		Banners:             input.Banners,
	}
}

func ModelRegistToCore(data []database.Jobseeker) []jobseeker.JobseekerRegistCore {
	var datasRegist []jobseeker.JobseekerRegistCore
	for _, input := range data {
		var inputData = jobseeker.JobseekerRegistCore{
			Full_name: input.Full_name,
			Username:  input.Username,
			Email:     input.Email,
			Password:  input.Password,
		}
		datasRegist = append(datasRegist, inputData)
	}

	return datasRegist
}

func ModelCVToCore(input database.CV) jobseeker.CVCore {
	return jobseeker.CVCore{
		ID:          input.ID,
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

func ModelEduToCore(input database.Education) jobseeker.EducationCore {
	return jobseeker.EducationCore{
		ID:              input.ID,
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}

func ModelEdusToCore(data []database.Education) []jobseeker.EducationCore {
	var edusData []jobseeker.EducationCore
	for _, input := range data {
		var eduInput = jobseeker.EducationCore{
			ID:              input.ID,
			JobseekerID:     input.JobseekerID,
			Education_level: input.Education_level,
			Major:           input.Major,
			Graduation_date: input.Graduation_date,
		}
		edusData = append(edusData, eduInput)
	}

	return edusData
}

func ModelLicenseToCore(input database.License) jobseeker.LicenseCore {
	return jobseeker.LicenseCore{
		ID:             input.ID,
		JobseekerID:    input.JobseekerID,
		License_name:   input.License_name,
		Published_date: input.Published_date,
		Expiry_date:    input.Expiry_date,
		License_file:   input.License_file,
	}
}

func ModelLicensesToCore(data []database.License) []jobseeker.LicenseCore {
	var licensesData []jobseeker.LicenseCore
	for _, input := range data {
		var licenseInput = jobseeker.LicenseCore{
			ID:             input.ID,
			JobseekerID:    input.JobseekerID,
			License_name:   input.License_name,
			Published_date: input.Published_date,
			Expiry_date:    input.Expiry_date,
			License_file:   input.License_file,
		}
		licensesData = append(licensesData, licenseInput)
	}

	return licensesData
}

func ModelSkillToCore(input database.Skill) jobseeker.SkillCore {
	return jobseeker.SkillCore{
		ID:          input.ID,
		JobseekerID: input.JobseekerID,
		Skill:       input.Skill,
		Description: input.Description,
	}
}

func ModelSkillsToCore(data []database.Skill) []jobseeker.SkillCore {
	var skillsData []jobseeker.SkillCore
	for _, input := range data {
		var skillInput = jobseeker.SkillCore{
			ID:          input.ID,
			JobseekerID: input.JobseekerID,
			Skill:       input.Skill,
			Description: input.Description,
		}
		skillsData = append(skillsData, skillInput)
	}
	return skillsData
}
