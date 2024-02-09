package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"errors"
)

type JobseekerService struct {
	jobseekerData jobseeker.JobseekerDataInterface
}

// dependency injection
func New(repo jobseeker.JobseekerDataInterface) jobseeker.JobseekerServiceInterface {
	return &JobseekerService{
		jobseekerData: repo,
	}
}

func (service *JobseekerService) Register(input jobseeker.JobseekerCore) error {
	// logic validation
	err := service.jobseekerData.Register(input)
	return err
}

func (service *JobseekerService) Login(email string, password string) (jobseeker.JobseekerCore, string, error) {
	if email == "" {
		return jobseeker.JobseekerCore{}, "", errors.New("email is required")
	} else if password == "" {
		return jobseeker.JobseekerCore{}, "", errors.New("password is required")
	}

	// get data from database that matches the given email
	resLogin, err := service.jobseekerData.Login(email)
	if err != nil {
		return jobseeker.JobseekerCore{}, "", errors.New(err.Error() + "login error, cannot retrieve data")
	}

	// checking given password
	checkPassword := responses.ComparePassword(password, resLogin.Password)
	if !checkPassword {
		return jobseeker.JobseekerCore{}, "", errors.New("login failed, wrong password")
	}

	// create token used for access other endpoints
	token, err := middlewares.CreateToken(resLogin.ID)
	if err != nil {
		return jobseeker.JobseekerCore{}, "", errors.New(err.Error() + "cannot create token")
	}
	return resLogin, token, nil
}

func (service *JobseekerService) UpdateProfile(userID uint, data jobseeker.JobseekerCore) error {
	// logic validation
	err := service.jobseekerData.UpdateProfile(userID, data)
	return err
}

// func (service *UserService) AddCareer(input user.CareerCore) error {
// 	// logic validation
// 	err := service.userData.AddCareer(input)
// 	return err
// }
