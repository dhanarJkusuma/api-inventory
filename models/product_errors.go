package models

import "errors"

var (
	ERR_PRODUCT_CONFLICT_SKU = errors.New("Attribute sku must be unique")
	ERR_PRODUCT_NOT_FOUND    = errors.New("Product not found")
	ERR_PRODUCT_DB           = errors.New("Internal server error")
)
