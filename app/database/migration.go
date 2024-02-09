package database

import (
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&Jobseeker{})
	db.AutoMigrate(&Career{})
	// db.AutoMigrate(&Store{})
	// db.AutoMigrate(&ShoppingCart{})
	// db.AutoMigrate(&ShoppingCartItem{})
	// db.AutoMigrate(&Order{})
	// db.AutoMigrate(&OrderItem{})
	// db.AutoMigrate(&Admin{})
	// db.AutoMigrate(&Payment{})
}
