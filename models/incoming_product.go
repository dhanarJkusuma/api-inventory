package models

import (
	"time"
)

type IncomingProduct struct {
	ID            int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Product       Product   `json:"product" validate:"-"`
	ProductID     int64     `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE" json:"productId"`
	ProductSKU    string    `gorm:"-" json:"productSKU" validate:"required"`
	TotalOrder    int64     `json:"totalOrder" validate:"required"`
	Total         int64     `json:"total" validate:"required"`
	Note          string    `json:"note"`
	BuyPrice      float64   `json:"buyPrice" validate:"required"`
	TotalAmount   float64   `json:"totalAmout"`
	NumberReceipt string    `gorm:"size:255;not null" json:"noReceipt"`
	DateFormatted string    `json:"dateFormatted" validate:"required"`
	Date          time.Time `json:"date"`
	LastUpdated   time.Time `json:"lastUpdated"`
	InsertedAt    time.Time `json:"insertedAt"`
}
