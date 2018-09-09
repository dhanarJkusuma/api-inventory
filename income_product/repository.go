package income_product

import (
	"inventory_app/models"
	"time"
)

type IncomeProductRepository interface {
	CreateNewIncomeProduct(p *models.IncomingProduct) (*models.IncomingProduct, error)
	FetchIncomeProduct(date time.Time, page int32, size int32) ([]models.IncomingProduct, error)
	GetDetailIncomeProductt(id int64) (*models.IncomingProduct, error)
	UpdateIncomeProduct(id int64, p models.IncomingProduct) (*models.IncomingProduct, error)
	DeleteIncomeProduct(id int64) error
}
