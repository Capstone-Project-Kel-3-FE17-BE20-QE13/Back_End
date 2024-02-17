package handler

import (
	"JobHuntz/features/jobseeker"
)

type JobseekerRequest struct {
	Full_name           string `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string `gorm:"not null;unique" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	Password            string `gorm:"not null" json:"password" form:"password"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Birth_date          string `json:"birth_date" form:"birth_date"`
	Gender              string `json:"gender" form:"gender"`
	Resume              string `json:"resume" form:"resume"`
	Status_Verification string `json:"stat_verif" form:"stat_verif"`
	Banners             string `json:"banners" form:"banners"`
}

type JobseekerRegistRequest struct {
	Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Username  string `gorm:"not null;unique" json:"username" form:"username"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Password  string `gorm:"not null" json:"password" form:"password"`
}

type JobseekerUpdateRequest struct {
	Full_name string `json:"full_name" form:"full_name"`
	Username  string `json:"username" form:"username"`
	//Email     string `json:"email" form:"email"`
	//Password   string `json:"password" form:"password"`
	Address    string `json:"address" form:"address"`
	Phone      string `json:"phone" form:"phone"`
	Birth_date string `json:"birth_date" form:"birth_date"`
	Gender     string `json:"gender" form:"gender"`
	Resume     string `json:"resume" form:"resume"`
	Banners    string `json:"banners" form:"banners"`
}

type CVRequest struct {
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CV_file     string `json:"cv_file" form:"cv_file"`
}

type CareerRequest struct {
	JobseekerID  uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
	Date_start   string `json:"date_start" form:"date_start"`
	Date_end     string `json:"date_end" form:"date_end"`
}

type EducationRequest struct {
	JobseekerID     uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string `json:"ed_level" form:"ed_level"`
	Major           string `json:"major" form:"major"`
	Graduation_date string `json:"grad_date" form:"grad_date"`
}

type LicenseRequest struct {
	JobseekerID    uint   `json:"jobseeker_id" form:"jobseeker_id"`
	License_name   string `json:"license_name" form:"license_name"`
	Published_date string `json:"pub_date" form:"pub_date"`
	Expiry_date    string `json:"exp_date" form:"exp_date"`
	License_file   string `json:"license" form:"license"`
}

type SkillRequest struct {
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Skill       string `json:"skill" form:"skill"`
	Description string `json:"description" form:"description"`
}

func RequestJobseekerToCore(input JobseekerRequest) jobseeker.JobseekerCore {
	return jobseeker.JobseekerCore{
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

func RequestJobseekerRegistToCore(input JobseekerRegistRequest) jobseeker.JobseekerRegistCore {
	return jobseeker.JobseekerRegistCore{
		Full_name: input.Full_name,
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
	}
}

func RequestJobseekerUpdateToCore(input JobseekerUpdateRequest) jobseeker.JobseekerUpdateCore {
	return jobseeker.JobseekerUpdateCore{
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

func RequestCVToCore(input CVRequest) jobseeker.CVCore {
	return jobseeker.CVCore{
		JobseekerID: input.JobseekerID,
		CV_file:     input.CV_file,
	}
}

func RequestCareerToCore(input CareerRequest) jobseeker.CareerCore {
	return jobseeker.CareerCore{
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}

func RequestEduToCore(input EducationRequest) jobseeker.EducationCore {
	return jobseeker.EducationCore{
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}

func RequestLicenseToCore(input LicenseRequest) jobseeker.LicenseCore {
	return jobseeker.LicenseCore{
		JobseekerID:    input.JobseekerID,
		License_name:   input.License_name,
		Published_date: input.Published_date,
		Expiry_date:    input.Expiry_date,
		License_file:   input.License_file,
	}
}

func RequestSkillToCore(input SkillRequest) jobseeker.SkillCore {
	return jobseeker.SkillCore{
		JobseekerID: input.JobseekerID,
		Skill:       input.Skill,
		Description: input.Description,
	}
}
