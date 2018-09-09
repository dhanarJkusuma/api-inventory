package models

import (
	"time"
)

type Product struct {
	ID          int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Sku         string `gorm:"size:255;unique;not null"`
	Name        string `gorm:"size:100"`
	Total       int64
	SellPrice   float64
	LastUpdated *time.Time
}
