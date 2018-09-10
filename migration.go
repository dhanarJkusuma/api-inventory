package main

import (
	"fmt"
	"inventory_app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func MigrateDatabase(db *gorm.DB) {
	fmt.Println("[Inventory App] Auto Migration Database")
	db.AutoMigrate(
		&models.Product{},
		&models.IncomingProduct{},
		&models.OutcomingProduct{},
	)
}
