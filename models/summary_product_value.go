package models

import (
	"time"
)

type ProductValue struct {
	ProductID     int64     `json:"productId"`
	ProductSKU    string    `json:"productSKU"`
	ProductName   string    `json:"productName"`
	Total         int64     `json:"total"`
	TotalBuyPrice float64   `json:"totalBuyPrice"`
	AvgBuyPrice   float64   `json:"avgBuyPrice"`
	TotalAmount   float64   `json:"totalAmount"`
	Date          time.Time `json:"date"`
}
