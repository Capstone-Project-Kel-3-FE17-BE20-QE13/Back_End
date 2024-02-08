package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/user"
)

func CoreUserToModel(input user.Core) database.User {
	return database.User{
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

func CoreCareerToModel(input user.CareerCore) database.Career {
	return database.Career{
		UserID:       input.UserID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}
