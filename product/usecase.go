package product

import (
	"inventory_app/models"
)

type ProductUsecase interface {
	CreateNewProduct(p *models.Product) (*models.Product, error)
	FetchProduct(page int, size int) ([]models.Product, error)
	GetProduct(id int64) (*models.Product, error)
	GetProductBySKU(sku string) (*models.Product, error)
	UpdateProduct(id int64, p *models.Product) (*models.Product, error)
	UpdateProductBySKU(sku string, p *models.Product) (*models.Product, error)
	DeleteProduct(id int64) error
	DeleteProductBySKU(sku string) error
}
