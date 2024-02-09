package jobseeker

import (
	"time"
)

type JobseekerCore struct {
	ID                  uint      `json:"id" form:"id"`
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
}

type CareerCore struct {
	ID           uint      `json:"id" form:"id"`
	JobseekerID  uint      `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

// interface untuk Service Layer
type JobseekerServiceInterface interface {
	Register(input JobseekerCore) error
	// AddCareer(input CareerCore) error
	// Login(email string, password string) (JobseekerCore, string, error)
}

// interface untuk Data Layer
type JobseekerDataInterface interface {
	Register(input JobseekerCore) error
	// AddCareer(input CareerCore) error
	// Login(email string) (JobseekerCore, error)
}
