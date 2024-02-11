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
	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Vacancy{})
	db.AutoMigrate(&Skill{})
}
