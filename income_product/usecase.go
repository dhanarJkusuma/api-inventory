package income_product

import (
	"inventory_app/models"
)

type IncomeProductUsecase interface {
	CreateNewIncomeProduct(p *models.IncomingProduct) (*models.IncomingProduct, error)
	CreateNewIncomeProductSilent(p *models.IncomingProduct) error
	FetchIncomeProduct(from string, page int, size int) ([]models.IncomingProduct, error)
	GetDetailIncomeProduct(id int64) (*models.IncomingProduct, error)
	GetDetailIncomeProductByNoReceipt(no string) (*models.IncomingProduct, error)
	UpdateIncomeProduct(id int64, p *models.IncomingProduct) (*models.IncomingProduct, error)
	DeleteIncomeProduct(id int64) error
	GetSummaryProductValue(from string, page int, size int) ([]models.ProductValue, error)
	BatchInsert(ops []models.IncomingProduct) error
}
