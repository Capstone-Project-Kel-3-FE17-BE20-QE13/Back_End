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

type Company struct {
	gorm.Model
	Company_name        string `gorm:"not null" json:"company_name" form:"company_name"`
	Full_name           string `json:"full_name" form:"full_name"`
	Email               string `gorm:"not null" json:"email" form:"email"`
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

type Job struct {
	gorm.Model
	Name        string
	TypeJob     string
	Salary      string
	Category    string
	JobDesc     string
	Requirement string
	CreatedBy   uint
	CompanyId   uint
}

type Apply struct {
	gorm.Model
	JobseekerID uint   `json:"jobseeker_id"`
	JobId       uint   `json:"job_id"`
	Status      string `json:"status"`
}
