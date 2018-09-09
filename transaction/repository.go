package transaction

import (
	"inventory_app/models"
	"time"
)

type TransactionRepository interface {
	CreateNewTransaction(t *models.Transaction) (*models.Transaction, error)
	FetchTransaction(from *time.Time, page int32, size int32) ([]models.Transaction, error)
	DetailTransaction(id int64) (*models.Transaction, error)
}
