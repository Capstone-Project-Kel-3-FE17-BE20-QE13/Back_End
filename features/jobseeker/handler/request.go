package handler

import (
	"JobHuntz/features/jobseeker"
	"time"
)

type JobseekerRequest struct {
	Full_name           string    `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string    `gorm:"not null" json:"username" form:"username"`
	Email               string    `gorm:"not null;unique" json:"email" form:"email"`
	Password            string    `gorm:"not null" json:"password" form:"password"`
	Address             string    `json:"address" form:"address"`
	Phone               string    `json:"phone" form:"phone"`
	Birth_date          time.Time `json:"birth_date" form:"birth_date"`
	Gender              string    `json:"gender" form:"gender"`
	Resume              string    `json:"resume" form:"resume"`
	Status_Verification string    `json:"stat_verif" form:"stat_verif"`
}

type CVRequest struct {
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CV_file     string `json:"cv_file" form:"cv_file"`
}

type CareerRequest struct {
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
}

type EducationRequest struct {
	JobseekerID     uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string    `json:"ed_level" form:"ed_level"`
	Major           string    `json:"major" form:"major"`
	Graduation_date time.Time `json:"grad_date" form:"grad_date"`
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

func RequestEduToCareer(input EducationRequest) jobseeker.EducationCore {
	return jobseeker.EducationCore{
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}
