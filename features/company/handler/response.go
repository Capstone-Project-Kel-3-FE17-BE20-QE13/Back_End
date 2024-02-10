package handler

import "JobHuntz/features/company"

type LoginResponse struct {
	// ID        uint   `json:"id" form:"id"`
	// Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Email string `gorm:"not null;unique" json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

func ResponCompanyToResponseLogin(input company.CompanyCore, token string) LoginResponse {
	return LoginResponse{
		// ID:        input.ID,
		// Full_name: input.Full_name,
		Email: input.Email,
		Token: token,
	}
}
