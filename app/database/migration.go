package database

import (
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&Jobseeker{})
	db.AutoMigrate(&Career{})
	db.AutoMigrate(&Company{})
}
