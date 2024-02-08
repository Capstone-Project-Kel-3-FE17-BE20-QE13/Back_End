package handler

import (
	"JobHuntz/features/user"
	"time"
)

type UserResponse struct {
	ID                  uint      `json:"id" form:"id"`
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
	CV                  []byte    `json:"cv" form:"cv"`
}

func CoreToResponse(input user.Core) UserResponse {
	return UserResponse{
		ID:                  input.ID,
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
