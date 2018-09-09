package repository

import (
	"fmt"
	"inventory_app/models"
	"inventory_app/outcome_product"
	"time"

	"github.com/jinzhu/gorm"
)

type outcomeProductRepository struct {
	DB *gorm.DB
}

func NewOutcomeProductRepository(db *gorm.DB) outcome_product.OutcomeProductRepository {
	return &outcomeProductRepository{
		DB: db,
	}
}

func (or *outcomeProductRepository) CreateNewOutcomeProduct(op *models.OutcomingProduct) (*models.OutcomingProduct, error) {
	op.LastUpdated = time.Now()
	op.InsertedAt = time.Now()
	op.TotalAmount = float64(op.Total) * float64(op.SellPrice)
	if dbc := or.DB.Create(op); dbc.Error != nil {
		return nil, models.ERR_RECORD_DB
	}
	or.DB.Model(&op).Related(&op.Product)
	return op, nil
}

func (or *outcomeProductRepository) FetchOutcomeProduct(from time.Time, page int, size int) ([]models.OutcomingProduct, error) {
	var outcomeProducts []models.OutcomingProduct
	offset := page * size
	if err := or.DB.Where("date >= ? ", from).Order("date desc").Offset(offset).Limit(size).Find(&outcomeProducts).Error; err != nil {
		fmt.Println("error")
		return []models.OutcomingProduct{}, models.ERR_RECORD_DB
	}
	for index := range outcomeProducts {
		or.DB.Model(outcomeProducts[index]).Related(&outcomeProducts[index].Product)
		outcomeProducts[index].ProductSKU = outcomeProducts[index].Product.Sku
	}
	return outcomeProducts, nil
}

func (or *outcomeProductRepository) GetDetailOutcomeProduct(id int64) (*models.OutcomingProduct, error) {
	var outcomeProduct models.OutcomingProduct
	if err := or.DB.Where("id = ?", id).First(&outcomeProduct).Error; err != nil {
		return nil, models.ERR_RECORD_NOT_FOUND
	}
	or.DB.Model(&outcomeProduct).Related(&outcomeProduct.Product)
	return &outcomeProduct, nil
}

func (or *outcomeProductRepository) UpdateOutcomeProduct(id int64, p models.OutcomingProduct) (*models.OutcomingProduct, models.OutcomingProduct, error) {
	var existOutProd models.OutcomingProduct
	if err := or.DB.Where("id = ?", id).First(&existOutProd).Error; err != nil {
		return nil, models.OutcomingProduct{}, models.ERR_RECORD_NOT_FOUND
	}
	or.DB.Model(&existOutProd).Related(&existOutProd.Product)
	oldData := existOutProd
	existOutProd.ProductID = p.ProductID
	existOutProd.ProductSKU = p.ProductSKU
	existOutProd.Product = p.Product
	existOutProd.Total = p.Total
	existOutProd.SellPrice = p.SellPrice
	existOutProd.Note = p.Note
	existOutProd.LastUpdated = time.Now()
	existOutProd.TotalAmount = float64(existOutProd.Total) * float64(existOutProd.SellPrice)
	existOutProd.Order = p.Order
	if err := or.DB.Save(&existOutProd).Error; err != nil {
		return nil, models.OutcomingProduct{}, models.ERR_RECORD_DB
	}
	return &existOutProd, oldData, nil
}

func (or *outcomeProductRepository) DeleteOutcomeProduct(id int64) error {
	if err := or.DB.Where("id = ?", id).Delete(&models.OutcomingProduct{}).Error; err != nil {
		return models.ERR_RECORD_DB
	}
	return nil
}
