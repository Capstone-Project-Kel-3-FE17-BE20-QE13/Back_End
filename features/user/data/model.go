package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/user"
)

func CoreToModel(input user.Core) database.User {
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
