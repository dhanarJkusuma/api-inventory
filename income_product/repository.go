package income_product

import (
	"inventory_app/models"
	"time"
)

type IncomeProductRepository interface {
	CreateNewIncomeProduct(p *models.IncomingProduct) (*models.IncomingProduct, error)
	FetchIncomeProduct(date time.Time, page int, size int) ([]models.IncomingProduct, error)
	GetDetailIncomeProduct(id int64) (*models.IncomingProduct, error)
	UpdateIncomeProduct(id int64, p models.IncomingProduct) (*models.IncomingProduct, models.IncomingProduct, error)
	DeleteIncomeProduct(id int64) error
}
