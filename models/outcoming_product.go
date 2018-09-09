package models

import (
	"time"
)

type OutcomingProduct struct {
	ID            int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Product       Product
	ProductID     int64     `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE"`
	ProductSKU    string    `gorm:"-" json:"productSKU"`
	Total         int64     `json:"total"`
	SellPrice     float64   `json:"sellPrice"`
	TotalAmount   float64   `json:"totalAmount"`
	Note          string    `json:"note"`
	DateFormatted string    `json:"dateFormatted"`
	Date          time.Time `json:"date"`
	LastUpdated   time.Time `json:"lastUpdated"`
	InsertedAt    time.Time `json:"insertedAt"`
	Order         int       `json:"-"`
	OrderID       string    `json:"orderId"`
	IsOrder       bool      `gorm:"-" json:"isOrder"`
}
