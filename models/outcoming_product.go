package models

import (
	"time"
)

type OutcomingProduct struct {
	ID            int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Product       Product   `validate:"-"`
	ProductID     int64     `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE"`
	ProductSKU    string    `gorm:"-" json:"productSKU" validate:"required"`
	Total         int64     `json:"total" validate:"required"`
	SellPrice     float64   `json:"sellPrice" validate:"required"`
	TotalAmount   float64   `json:"totalAmount"`
	Note          string    `json:"note"`
	DateFormatted string    `json:"dateFormatted" validate:"required"`
	Date          time.Time `json:"date"`
	LastUpdated   time.Time `json:"lastUpdated"`
	InsertedAt    time.Time `json:"insertedAt"`
	OrderID       string    `json:"orderId"`
}
