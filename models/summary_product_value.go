package models

import (
	"time"
)

type ProductValue struct {
	ProductID     int64
	ProductSKU    string
	ProductName   string
	Total         int64
	TotalBuyPrice float64
	AvgBuyPrice   float64
	TotalAmount   float64
	Date          time.Time
}
