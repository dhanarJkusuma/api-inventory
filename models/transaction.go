package models

import (
	"time"
)

type Transaction struct {
	ID           int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ProductName  string `gorm:"size:100"`
	ProductPrice float64
	Qty          int64
	TotalAmount  float64
	CreatedAt    *time.Time
}
