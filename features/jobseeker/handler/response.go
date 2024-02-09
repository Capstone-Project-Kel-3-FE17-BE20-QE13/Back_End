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

// func CoreSeekerToResponseLogin(input user.UserCore, token string) LoginResponse {
// 	return LoginResponse{
// 		ID:        input.ID,
// 		Full_name: input.Full_name,
// 		Email:     input.Email,
// 		Role:      input.Role,
// 		Token:     token,
// 	}
// }

// func CoreCareerToResponse(input user.CareerCore) CareerResponse {
// 	return CareerResponse{
// 		ID:           input.ID,
// 		JobseekerID:  input.JobseekerID,
// 		Position:     input.Position,
// 		Company_name: input.Company_name,
// 		Date_start:   input.Date_start,
// 		Date_end:     input.Date_end,
// 	}
// }
