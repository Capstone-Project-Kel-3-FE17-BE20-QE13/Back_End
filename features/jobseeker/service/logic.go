package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"database/sql"
	"errors"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func (service *JobseekerService) GetjobseekerByCompany(input uint) (*jobseeker.JobseekerCore, error) {
	// logic validation
	res, err := service.jobseekerData.GetByIdJobSeeker(input)
	return res, err
}

// GetByIdJobSeeker implements jobseeker.JobseekerServiceInterface.
func (service *JobseekerService) GetByIdJobSeeker(id uint) (*jobseeker.JobseekerCore, error) {
	result, err := service.jobseekerData.GetByIdJobSeeker(id)
	return result, err
}

func (service *JobseekerService) RegisterValidation(input jobseeker.JobseekerRegistCore) error {
	// logic validation
	if input.Email == "" {
		return errors.New("email is required")
	}

	if input.Full_name == "" || input.Password == "" || input.Username == "" {
		return errors.New("please complete your data")
	}

	// get data from database that matches the given email
	resRegist1, _ := service.jobseekerData.AllEmails(input.Email)
	if resRegist1.Email == input.Email {
		return errors.New("duplicate entry")
	}

	resRegist2, _ := service.jobseekerData.AllUsernames(input.Username)
	if resRegist2.Username == input.Username {
		return errors.New("duplicate entry")
	}

	if len(input.Password) < 8 {
		return errors.New("your password is not valid")
	}

	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	for _, v := range input.Full_name {
		if !strings.Contains(characters, string(v)) {
			return errors.New("your name is not valid")
		}
	}

	if strings.Contains(input.Password, " ") {
		return errors.New("your password is not valid")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regexpEmail := regexp.MustCompile(emailRegex)
	if !regexpEmail.MatchString(input.Email) {
		return errors.New("please enter a valid email")
	}

	return nil
}

func (service *JobseekerService) Register(input jobseeker.JobseekerRegistCore) error {
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

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Membuat objek regex dari ekspresi reguler
	regexpEmail := regexp.MustCompile(emailRegex)
	if !regexpEmail.MatchString(email) {
		return jobseeker.JobseekerCore{}, "", errors.New("please enter a valid email")
	}

	// get data from database that matches the given email
	resLogin, err := service.jobseekerData.AllEmails(email)
	if err != nil {
		return jobseeker.JobseekerCore{}, "", errors.New("wrong email")
	}

	// checking given password
	checkPassword := responses.ComparePassword(password, resLogin.Password)
	if !checkPassword {
		return jobseeker.JobseekerCore{}, "", errors.New("login failed, wrong password")
	}

	// create token used for access other endpoints
	token, err := middlewares.CreateToken(int(resLogin.ID))
	if err != nil {
		return jobseeker.JobseekerCore{}, "", errors.New(err.Error() + "cannot create token")
	}
	return resLogin, token, nil
}

func (service *JobseekerService) UpdateValidation(input jobseeker.JobseekerUpdateCore) error {
	if input.Username != "" {
		resData2, _ := service.jobseekerData.AllUsernames(input.Username)
		if resData2.Username == input.Username {
			return errors.New("username is already used")
		}
	}

	if input.Phone != "" {
		numbers := "1234567890"
		for _, v := range input.Phone {
			if !strings.Contains(numbers, string(v)) {
				return errors.New("please enter a valid phone number")
			}
		}
	}

	return nil
}

func (service *JobseekerService) UpdateProfile(seekerID uint, data jobseeker.JobseekerUpdateCore) error {
	// logic validation
	err := service.jobseekerData.UpdateProfile(seekerID, data)
	return err
}

func (service *JobseekerService) PDF(input *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	// logic validation
	res, err := service.jobseekerData.PDF(input)
	return res, err
}

func (service *JobseekerService) Photo(input *multipart.FileHeader) (*uploader.UploadResult, error) {
	// logic validation
	res, err := service.jobseekerData.Photo(input)
	return res, err
}

func (service *JobseekerService) CountCV(dbRaw *sql.DB, seekerID uint) (uint, error) {
	// logic validation
	res, err := service.jobseekerData.CountCV(dbRaw, seekerID)
	return res, err
}

func (service *JobseekerService) AddCV(input jobseeker.CVCore, count uint) error {
	if count == 1 {
		return errors.New("your cv is already exist")
	}

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
