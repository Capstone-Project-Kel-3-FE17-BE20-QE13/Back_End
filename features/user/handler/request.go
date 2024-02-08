package handler

import (
	"JobHuntz/features/user"
	"time"
)

type UserRequest struct {
	Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Password  string `gorm:"not null" json:"password" form:"password"`
	Role      string `json:"role" form:"role"`
}

type JobseekerRequest struct {
	Full_name           string    `gorm:"not null" json:"full_name" form:"full_name"`
	Email               string    `gorm:"not null;unique" json:"email" form:"email"`
	Password            string    `gorm:"not null" json:"password" form:"password"`
	Role                string    `json:"role" form:"role"`
	Username            string    `gorm:"not null" json:"username" form:"username"`
	Address             string    `json:"alamat" form:"alamat"`
	Phone               string    `json:"phone" form:"phone"`
	Status_Verification string    `json:"stat_verif" form:"stat_verif"`
	Birth_date          time.Time `json:"birth_date" form:"birth_date"`
	Gender              string    `json:"gender" form:"gender"`
	Resume              string    `json:"resume" form:"resume"`
	CV                  string    `json:"cv" form:"cv"`
}

type CareerRequest struct {
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
}

func RequestUserToCore(input UserRequest) user.UserCore {
	return user.UserCore{
		Full_name: input.Full_name,
		Email:     input.Email,
		Password:  input.Password,
		Role:      input.Role,
	}
}

func RequestSeekerToCore(input JobseekerRequest) user.JobseekerCore {
	return user.JobseekerCore{
		Full_name:           input.Full_name,
		Email:               input.Email,
		Password:            input.Password,
		Role:                input.Role,
		Username:            input.Username,
		Address:             input.Address,
		Phone:               input.Phone,
		Status_Verification: input.Status_Verification,
		Birth_date:          input.Birth_date,
		Gender:              input.Gender,
		Resume:              input.Resume,
		CV:                  input.CV,
	}
}

func RequestCareerToCore(input CareerRequest) user.CareerCore {
	return user.CareerCore{
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}
