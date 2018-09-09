package models

import (
	"time"
)

type Product struct {
	ID          int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Sku         string    `gorm:"size:255;unique;not null" json:"sku"`
	Name        string    `gorm:"size:100" json:"name"`
	Total       int64     `json:"total"`
	SellPrice   float64   `json:"sellPrice"`
	LastUpdated time.Time `json:"lastUpdated"`
}
