package handler

import (
	"JobHuntz/features/jobseeker"
	"time"
)

type LoginResponse struct {
	ID        uint   `json:"id" form:"id"`
	Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Token     string `json:"token" form:"token"`
}

type JobseekerResponse struct {
	ID                  uint      `json:"id" form:"id"`
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

type CareerResponse struct {
	ID           uint      `json:"id" form:"id"`
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
}

type EducationResponse struct {
	ID              uint      `json:"id" form:"id"`
	JobseekerID     uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string    `json:"ed_level" form:"ed_level"`
	Major           string    `json:"major" form:"major"`
	Graduation_date time.Time `json:"grad_date" form:"grad_date"`
}

func CoreJobseekerToResponse(input jobseeker.JobseekerCore) JobseekerResponse {
	return JobseekerResponse{
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

func CoreJobseekerToResponseLogin(input jobseeker.JobseekerCore, token string) LoginResponse {
	return LoginResponse{
		ID:        input.ID,
		Full_name: input.Full_name,
		Email:     input.Email,
		Token:     token,
	}
}

func CoreCareerToResponse(input jobseeker.CareerCore) CareerResponse {
	return CareerResponse{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}

func CoreCareersToResponse(input []jobseeker.CareerCore) []CareerResponse {
	var careerResponses []CareerResponse
	for _, v := range input {
		var careerInput = CareerResponse{
			ID:           v.ID,
			JobseekerID:  v.JobseekerID,
			Position:     v.Position,
			Company_name: v.Company_name,
			Date_start:   v.Date_start,
			Date_end:     v.Date_end,
		}
		careerResponses = append(careerResponses, careerInput)
	}

	return careerResponses
}

func CoreEducationToResponse(input jobseeker.EducationCore) EducationResponse {
	return EducationResponse{
		ID:              input.ID,
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}

func CoreEdusToResponse(input []jobseeker.EducationCore) []EducationResponse {
	var eduResponses []EducationResponse
	for _, input := range input {
		var eduInput = EducationResponse{
			ID:              input.ID,
			JobseekerID:     input.JobseekerID,
			Education_level: input.Education_level,
			Major:           input.Major,
			Graduation_date: input.Graduation_date,
		}
		eduResponses = append(eduResponses, eduInput)
	}

	return eduResponses
}
