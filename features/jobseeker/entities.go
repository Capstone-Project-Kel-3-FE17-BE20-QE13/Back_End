package jobseeker

import (
	"database/sql"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type JobseekerCore struct {
	ID                  uint   `json:"id" form:"id"`
	Full_name           string `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string `gorm:"not null:unique" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	Password            string `gorm:"not null" json:"password" form:"password"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Birth_date          string `json:"birth_date" form:"birth_date"`
	Gender              string `json:"gender" form:"gender"`
	Resume              string `json:"resume" form:"resume"`
	Status_Verification string `json:"stat_verif" form:"stat_verif"`
	Banners             string `json:"banners" form:"banners"`
	Roles               string
	Careers             []CareerCore
	Educations          []EducationCore
	Cvs                 []CVCore
	Licenses            []LicenseCore
	Skills              []SkillCore
}

type JobseekerRegistCore struct {
	Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Username  string `gorm:"not null" json:"username" form:"username"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Password  string `gorm:"not null" json:"password" form:"password"`
}

type JobseekerUpdateCore struct {
	Full_name  string `json:"full_name" form:"full_name"`
	Username   string `json:"username" form:"username"`
	Address    string `json:"address" form:"address"`
	Phone      string `json:"phone" form:"phone"`
	Birth_date string `json:"birth_date" form:"birth_date"`
	Gender     string `json:"gender" form:"gender"`
	Resume     string `json:"resume" form:"resume"`
	Banners    string `json:"banners" form:"banners"`
}

type CVCore struct {
	ID          uint   `json:"id" form:"id"`
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CV_file     string `json:"cv_file" form:"cv_file"`
	Jobseeker   *JobseekerCore
}

type CareerCore struct {
	ID           uint      `json:"id" form:"id"`
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   string    `json:"date_start" form:"date_start"`
	Date_end     string    `json:"date_end" form:"date_end"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
	Jobseeker    *JobseekerCore
}

type EducationCore struct {
	ID              uint   `json:"id" form:"id"`
	JobseekerID     uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string `json:"ed_level" form:"ed_level"`
	Major           string `json:"major" form:"major"`
	Graduation_date string `json:"grad_date" form:"grad_date"`
	Jobseeker       *JobseekerCore
}

type LicenseCore struct {
	ID             uint   `json:"id" form:"id"`
	JobseekerID    uint   `json:"jobseeker_id" form:"jobseeker_id"`
	License_name   string `json:"license_name" form:"license_name"`
	Published_date string `json:"pub_date" form:"pub_date"`
	Expiry_date    string `json:"exp_date" form:"exp_date"`
	License_file   string `json:"license" form:"license"`
	Jobseeker      *JobseekerCore
}

type SkillCore struct {
	ID          uint   `json:"id" form:"id"`
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Skill       string `json:"skill" form:"skill"`
	Description string `json:"description" form:"description"`
	Jobseeker   *JobseekerCore
}

// interface untuk Service Layer
type JobseekerServiceInterface interface {
	RegisterValidation(input JobseekerRegistCore) error
	Register(input JobseekerRegistCore) error
	Login(email string, password string) (JobseekerCore, string, error)
	UpdateValidation(input JobseekerUpdateCore) error
	UpdateProfile(seekerID uint, data JobseekerUpdateCore) error
	PDF(input *multipart.FileHeader) (*s3manager.UploadOutput, error)
	Photo(input *multipart.FileHeader) (*uploader.UploadResult, error)
	CountCV(dbRaw *sql.DB, seekerID uint) (uint, error)
	AddCV(input CVCore, count uint) error
	ReadCV(seekerID uint) (CVCore, error)
	UpdateCV(input CVCore) error
	RemoveCV(seekerID uint) error
	AddCareer(input CareerCore) error
	GetCareerByID(career_id uint) (CareerCore, error)
	GetCareerList(seekerID uint) ([]CareerCore, error)
	UpdateCareer(careerID_int uint, input CareerCore) error
	RemoveCareer(career_id uint) error
	AddEducation(input EducationCore) error
	GetEduByID(eduID uint) (EducationCore, error)
	GetEduList(seekerID uint) ([]EducationCore, error)
	UpdateEducation(eduID uint, data EducationCore) error
	RemoveEducation(eduID uint) error
	AddLicense(input LicenseCore) error
	GetLicenseByID(licenseID uint) (LicenseCore, error)
	GetLicenseList(seekerID uint) ([]LicenseCore, error)
	UpdateLicense(licenseID uint, data LicenseCore) error
	RemoveLicense(licenseID uint) error
	AddSkill(input SkillCore) error
	GetSkillByID(skillID uint) (SkillCore, error)
	GetSkillList(seekerID uint) ([]SkillCore, error)
	UpdateSkill(skillID uint, data SkillCore) error
	RemoveSkill(skillID uint) error
	GetByIdJobSeeker(id uint) (*JobseekerCore, error)
	GetjobseekerByCompany(input uint) (*JobseekerCore, error)
}

// interface untuk Data Layer
type JobseekerDataInterface interface {
	Register(input JobseekerRegistCore) error
	AllEmails(email string) (JobseekerCore, error)
	AllUsernames(username string) (JobseekerCore, error)
	UpdateProfile(seekerID uint, data JobseekerUpdateCore) error
	PDF(input *multipart.FileHeader) (*s3manager.UploadOutput, error)
	Photo(input *multipart.FileHeader) (*uploader.UploadResult, error)
	CountCV(dbRaw *sql.DB, seekerID uint) (uint, error)
	AddCV(input CVCore) error
	ReadCV(seekerID uint) (CVCore, error)
	UpdateCV(input CVCore) error
	RemoveCV(seekerID uint) error
	AddCareer(input CareerCore) error
	GetCareerByID(career_id uint) (CareerCore, error)
	GetCareerList(seekerID uint) ([]CareerCore, error)
	UpdateCareer(careerID_int uint, input CareerCore) error
	RemoveCareer(career_id uint) error
	AddEducation(input EducationCore) error
	GetEduByID(eduID uint) (EducationCore, error)
	GetEduList(seekerID uint) ([]EducationCore, error)
	UpdateEducation(eduID uint, data EducationCore) error
	RemoveEducation(eduID uint) error
	AddLicense(input LicenseCore) error
	GetLicenseByID(licenseID uint) (LicenseCore, error)
	GetLicenseList(seekerID uint) ([]LicenseCore, error)
	UpdateLicense(licenseID uint, data LicenseCore) error
	RemoveLicense(licenseID uint) error
	AddSkill(input SkillCore) error
	GetSkillByID(skillID uint) (SkillCore, error)
	GetSkillList(seekerID uint) ([]SkillCore, error)
	UpdateSkill(skillID uint, data SkillCore) error
	RemoveSkill(skillID uint) error
	GetByIdJobSeeker(id uint) (*JobseekerCore, error)
	GetjobseekerByCompany(input uint) (*JobseekerCore, error)
}
