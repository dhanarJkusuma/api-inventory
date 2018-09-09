package models

import (
	"time"
)

type IncomingProduct struct {
	ID          int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Product     Product
	ProductID   int64 `sql:"type:bigint REFERENCES products(id) ON DELETE CASCADE"`
	Total       int64
	Note        string
	LastUpdated *time.Time
}
