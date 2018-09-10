package models

import (
	"time"
)

type Sales struct {
	OrderID     string    `json:"orderId"`
	Date        time.Time `json:"date"`
	ProductSKU  string    `json:"productSku"`
	ProductName string    `json:"productName"`
	Total       int64     `json:"total"`
	SellPrice   float64   `json:"sellPrice"`
	TotalAmount float64   `json:"totalAmount"`
	BuyPrice    float64   `json:"buyPrice"`
	Profit      float64   `json:"profit"`
}
