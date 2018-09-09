package models

import (
	"time"
)

type IncomingProduct struct {
	ID          int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Product     Product   `json:"product"`
	ProductID   int64     `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE" json:"productId"`
	ProductSKU  string    `gorm:"-" json:"productSKU"`
	Total       int64     `json:"total"`
	Note        string    `json:"note"`
	LastUpdated time.Time `json:"lastUpdated"`
}
