package models

import (
	"time"
)

type IncomingProduct struct {
	ID          int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Product     Product
	Total       int64
	Note        string
	LastUpdated *time.Time
}
