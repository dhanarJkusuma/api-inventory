package main

import (
	"fmt"
	"inventory_app/models"

	"github.com/jinzhu/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	fmt.Println("Migrating Database... ")
	db.AutoMigrate(
		&models.Product{},
		&models.IncomingProduct{},
		&models.OutcomingProduct{},
		&models.Transaction{},
	)
}
