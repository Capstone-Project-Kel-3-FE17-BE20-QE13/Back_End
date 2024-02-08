package service

import (
	"JobHuntz/features/user"
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

func (service *UserService) Register(input user.Core) error {
	// logic validation
	err := service.userData.Register(input)
	return err
}

func (service *UserService) AddCareer(input user.CareerCore) error {
	// logic validation
	err := service.userData.AddCareer(input)
	return err
}
