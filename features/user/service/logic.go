package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/user"
	"JobHuntz/utils/responses"
	"errors"
)

type UserService struct {
	userData user.UserDataInterface
}

// dependency injection
func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

func (service *UserService) Register(input user.UserCore) error {
	// logic validation
	err := service.userData.Register(input)
	return err
}

func (service *UserService) Login(email string, password string) (user.UserCore, string, error) {
	if email == "" {
		return user.UserCore{}, "", errors.New("email is required")
	} else if password == "" {
		return user.UserCore{}, "", errors.New("password is required")
	}

	// get data from database that matches the given email
	resLogin, err := service.userData.Login(email)
	if err != nil {
		return user.UserCore{}, "", errors.New(err.Error() + "login error, cannot retrieve data")
	}

	// checking given password
	checkPassword := responses.ComparePassword(password, resLogin.Password)
	if !checkPassword {
		return user.UserCore{}, "", errors.New("login failed, wrong password")
	}

	// create token used for access other endpoints
	token, err := middlewares.CreateToken(resLogin.ID)
	if err != nil {
		return user.UserCore{}, "", errors.New(err.Error() + "cannot create token")
	}
	return resLogin, token, nil
}

func (service *UserService) AddCareer(input user.CareerCore) error {
	// logic validation
	err := service.userData.AddCareer(input)
	return err
}
