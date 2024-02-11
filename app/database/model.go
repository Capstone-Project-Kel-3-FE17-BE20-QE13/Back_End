package database

import (
	"time"

	"gorm.io/gorm"
)

type Jobseeker struct {
	gorm.Model
	Full_name           string    `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string    `gorm:"not null" json:"username" form:"username"`
	Email               string    `gorm:"not null;unique" json:"email" form:"email"`
	Password            string    `gorm:"not null" json:"password" form:"password"`
	Address             string    `json:"address" form:"address"`
	Phone               string    `json:"phone" form:"phone"`
	Birth_date          time.Time `json:"birth_date" form:"birth_date"`
	Gender              string    `json:"gender" form:"gender"`
	Resume              string    `json:"resume" form:"resume"`
	Status_Verification string    `json:"stat_verif" form:"stat_verif"`
	CV                  CV
}

type CV struct {
	gorm.Model
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CV_file     string `json:"cv_file" form:"cv_file"`
}

type Career struct {
	gorm.Model
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
	Jobseeker    Jobseeker
}

type Education struct {
	gorm.Model
	JobseekerID     uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string    `json:"ed_level" form:"ed_level"`
	Major           string    `json:"major" form:"major"`
	Graduation_date time.Time `json:"grad_date" form:"grad_date"`
	Jobseeker       Jobseeker
}

type License struct {
	gorm.Model
	JobseekerID    uint      `json:"jobseeker_id" form:"jobseeker_id"`
	License_name   string    `json:"license_name" form:"license_name"`
	Published_date time.Time `json:"pub_date" form:"pub_date"`
	Expiry_date    time.Time `json:"exp_date" form:"exp_date"`
	License_file   string    `json:"license" form:"license"`
	Jobseeker      Jobseeker
}

type Skill struct {
	gorm.Model
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Skill       string `json:"skill" form:"skill"`
	Description string `json:"description" form:"description"`
	Jobseeker   Jobseeker
}

type Company struct {
	gorm.Model
	Company_name        string `gorm:"not null" json:"company_name" form:"company_name"`
	Full_name           string `json:"full_name" form:"full_name"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	Password            string `gorm:"not null" json:"password" form:"password"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Company_type        string `gorm:"not null" json:"company_type" form:"company_type"`
	Company_size        string `gorm:"not null" json:"company_size" form:"company_size"`
	Website             string `gorm:"not null" json:"website" form:"website"`
	Description         string `json:"description" form:"description"`
	Status_Verification string `json:"status_verification" form:"status_verification"`
	Banners             string `json:"banners" form:"banners"`
}

type Vacancy struct {
	gorm.Model
	CompanyID       uint   `json:"company_id" form:"company_id"`
	Name            string `json:"name" form:"name"`
	Job_type        string `json:"job_type" form:"job_type"`
	Salary_range    string `json:"salary_range" form:"salary_range"`
	Category        string `json:"category" form:"category"`
	Job_description string `json:"job_desc" form:"job_desc"`
	Job_requirement string `json:"job_req" form:"job_req"`
	Created_by      uint   `json:"created_by" form:"created_by"`
	Company         Company
}

type Application struct {
	gorm.Model
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Status_application string `json:"stat_app" form:"stat_app"`
	Jobseeker          Jobseeker
	Vacancy            Vacancy
}

type Favorites struct {
	gorm.Model
	JobseekerID uint   `json:"jobseeker_id"`
	JobId       uint   `json:"job_id"`
	Name        string `json:"name"`
}
