package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"errors"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

func (service *JobseekerService) UpdateProfile(seekerID uint, data jobseeker.JobseekerCore) error {
	// logic validation
	err := service.jobseekerData.UpdateProfile(seekerID, data)
	return err
}

func (service *JobseekerService) CV(input *multipart.FileHeader) (*uploader.UploadResult, error) {
	// logic validation
	res, err := service.jobseekerData.CV(input)
	return res, err
}

func (service *JobseekerService) AddCV(input jobseeker.CVCore) error {
	// logic validation
	err := service.jobseekerData.AddCV(input)
	return err
}

func (service *JobseekerService) ReadCV(seekerID uint) (jobseeker.CVCore, error) {
	// logic validation
	res, err := service.jobseekerData.ReadCV(seekerID)
	return res, err
}

func (service *JobseekerService) UpdateCV(input jobseeker.CVCore) error {
	// logic validation
	err := service.jobseekerData.UpdateCV(input)
	return err
}

// func (service *UserService) AddCareer(input user.CareerCore) error {
// 	// logic validation
// 	err := service.userData.AddCareer(input)
// 	return err
// }
