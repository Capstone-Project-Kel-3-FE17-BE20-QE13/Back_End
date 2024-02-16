package database

import (
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&Jobseeker{})
	db.AutoMigrate(&CV{})
	db.AutoMigrate(&Career{})
	db.AutoMigrate(&Education{})
	db.AutoMigrate(&License{})
	db.AutoMigrate(&Skill{})
	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Vacancy{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&Application{})
	db.AutoMigrate(&OrderCompany{})
	db.AutoMigrate(&OrderJobseeker{})
	db.AutoMigrate(&Payment{})
}
