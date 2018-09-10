package outcome_product

import (
	"inventory_app/models"
	"time"
)

type OutcomeProductRepository interface {
	CreateNewOutcomeProduct(op *models.OutcomingProduct) (*models.OutcomingProduct, error)
	FetchOutcomeProduct(from time.Time, page int, size int) ([]models.OutcomingProduct, error)
	GetDetailOutcomeProduct(id int64) (*models.OutcomingProduct, error)
	UpdateOutcomeProduct(id int64, p models.OutcomingProduct) (*models.OutcomingProduct, models.OutcomingProduct, error)
	DeleteOutcomeProduct(id int64) error
	GetSalesReport(startDate string, endDate string, page int, size int) (*models.SummarySales, error)
}
