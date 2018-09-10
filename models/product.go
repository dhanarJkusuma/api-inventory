package models

import (
	"errors"
	"time"
)

type Product struct {
	ID          int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Sku         string    `gorm:"size:255;unique;not null" json:"sku" validate:"required"`
	Name        string    `gorm:"size:100" json:"name" validate:"required"`
	Total       int64     `json:"total"`
	LastUpdated time.Time `json:"lastUpdated"`
}

var (
	ERR_PRODUCT_CONFLICT_SKU = errors.New("Attribute sku must be unique")
	ERR_PRODUCT_NOT_FOUND    = errors.New("Product not found")
	ERR_PRODUCT_DB           = errors.New("Internal server error")
)
