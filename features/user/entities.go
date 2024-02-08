package user

import (
	"time"
)

type Core struct {
	ID                  uint      `json:"id" form:"id"`
	Full_name           string    `gorm:"not null" json:"full_name" form:"full_name"`
	Email               string    `gorm:"not null;unique" json:"email" form:"email"`
	Password            string    `gorm:"not null" json:"password" form:"password"`
	Role                string    `json:"role" form:"role"`
	Username            string    `gorm:"not null" json:"username" form:"username"`
	Address             string    `json:"alamat" form:"alamat"`
	Phone               string    `json:"phone" form:"phone"`
	Status_Verification string    `json:"stat_verif" form:"stat_verif"`
	Birth_date          time.Time `json:"birth_date" form:"birth_date"`
	Gender              string    `json:"gender" form:"gender"`
	Resume              string    `json:"resume" form:"resume"`
	CV                  []byte    `json:"cv" form:"cv"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" form:"updated_at"`
}

type CareerCore struct {
	ID           uint      `json:"id" form:"id"`
	UserID       uint      `json:"user_id" form:"user_id"`
	Position     string    `json:"position" form:"position"`
	Company_name string    `json:"company_name" form:"company_name"`
	Date_start   time.Time `json:"date_start" form:"date_start"`
	Date_end     time.Time `json:"date_end" form:"date_end"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Register(input Core) error
	AddCareer(input CareerCore) error
}

// interface untuk Data Layer
type UserDataInterface interface {
	Register(input Core) error
	AddCareer(input CareerCore) error
}
