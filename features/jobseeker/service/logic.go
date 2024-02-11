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

func (service *JobseekerService) RemoveCV(input uint) error {
	// logic validation
	err := service.jobseekerData.RemoveCV(input)
	return err
}

func (service *JobseekerService) AddCareer(input jobseeker.CareerCore) error {
	// logic validation
	err := service.jobseekerData.AddCareer(input)
	return err
}

func (service *JobseekerService) GetCareerByID(input uint) (jobseeker.CareerCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetCareerByID(input)
	return res, err
}

func (service *JobseekerService) GetCareerList(input uint) ([]jobseeker.CareerCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetCareerList(input)
	return res, err
}

func (service *JobseekerService) UpdateCareer(career_id uint, input jobseeker.CareerCore) error {
	// logic validation
	err := service.jobseekerData.UpdateCareer(career_id, input)
	return err
}

func (service *JobseekerService) RemoveCareer(career_id uint) error {
	// logic validation
	err := service.jobseekerData.RemoveCareer(career_id)
	return err
}

func (service *JobseekerService) AddEducation(input jobseeker.EducationCore) error {
	// logic validation
	err := service.jobseekerData.AddEducation(input)
	return err
}

func (service *JobseekerService) GetEduByID(input uint) (jobseeker.EducationCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetEduByID(input)
	return res, err
}

func (service *JobseekerService) GetEduList(input uint) ([]jobseeker.EducationCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetEduList(input)
	return res, err
}

func (service *JobseekerService) UpdateEducation(eduID uint, input jobseeker.EducationCore) error {
	// logic validation
	err := service.jobseekerData.UpdateEducation(eduID, input)
	return err
}

func (service *JobseekerService) RemoveEducation(eduID uint) error {
	// logic validation
	err := service.jobseekerData.RemoveEducation(eduID)
	return err
}

func (service *JobseekerService) AddLicense(input jobseeker.LicenseCore) error {
	// logic validation
	err := service.jobseekerData.AddLicense(input)
	return err
}

func (service *JobseekerService) GetLicenseByID(input uint) (jobseeker.LicenseCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetLicenseByID(input)
	return res, err
}

func (service *JobseekerService) GetLicenseList(input uint) ([]jobseeker.LicenseCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetLicenseList(input)
	return res, err
}

func (service *JobseekerService) UpdateLicense(licenseID uint, input jobseeker.LicenseCore) error {
	// logic validation
	err := service.jobseekerData.UpdateLicense(licenseID, input)
	return err
}

func (service *JobseekerService) RemoveLicense(licenseID uint) error {
	// logic validation
	err := service.jobseekerData.RemoveLicense(licenseID)
	return err
}

func (service *JobseekerService) AddSkill(input jobseeker.SkillCore) error {
	// logic validation
	err := service.jobseekerData.AddSkill(input)
	return err
}

func (service *JobseekerService) GetSkillByID(input uint) (jobseeker.SkillCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetSkillByID(input)
	return res, err
}

func (service *JobseekerService) GetSkillList(input uint) ([]jobseeker.SkillCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetSkillList(input)
	return res, err
}

func (service *JobseekerService) UpdateSkill(skillID uint, input jobseeker.SkillCore) error {
	// logic validation
	err := service.jobseekerData.UpdateSkill(skillID, input)
	return err
}

func (service *JobseekerService) RemoveSkill(skillID uint) error {
	// logic validation
	err := service.jobseekerData.RemoveSkill(skillID)
	return err
}
