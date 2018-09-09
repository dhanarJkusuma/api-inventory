package outcome_product

import (
	"inventory_app/models"
	"time"
)

type OutcomeProductUsecase interface {
	CreateNewOutcomeProduct(p *models.OutcomingProduct) (*models.OutcomingProduct, error)
	FetchOutcomeProduct(date *time.Time, page int32, size int32) ([]models.OutcomingProduct, error)
	GetDetailOutcomeProductt(id int64) (*models.OutcomingProduct, error)
	UpdateOutcomeProduct(id int64, p *models.OutcomingProduct) (*models.OutcomingProduct, error)
	DeleteOutcomeProduct(id int64) error
}
