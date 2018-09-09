package models

import (
	"time"
)

type IncomingProduct struct {
	ID            int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Product       Product   `json:"product"`
	ProductID     int64     `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE" json:"productId"`
	ProductSKU    string    `gorm:"-" json:"productSKU"`
	TotalOrder    int64     `json:"totalOrder"`
	Total         int64     `json:"total"`
	Note          string    `json:"note"`
	BuyPrice      float64   `json:"buyPrice"`
	TotalAmount   float64   `json:"totalAmout"`
	NumberReceipt string    `gorm:"size:255;not null" json:"noReceipt"`
	DateFormatted string    `json:"dateFormatted"`
	Date          time.Time `json:"date"`
	LastUpdated   time.Time `json:"lastUpdated"`
	InsertedAt    time.Time `json:"insertedAt"`
}
