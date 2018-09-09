package product

import (
	"inventory_app/models"
)

type ProductRepository interface {
	CreateNewProduct(p *models.Product) (*models.Product, error)
	FetchProduct(page int32, size int32) ([]models.Product, error)
	GetDetailProduct(id int64) (*models.Product, error)
	UpdateProduct(id int64, p *models.Product) (*models.Product, error)
	DeleteProduct(id int64) error
}
