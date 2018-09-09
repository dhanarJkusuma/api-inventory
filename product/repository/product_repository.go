package repository

import (
	"inventory_app/models"
	"inventory_app/product"
	"time"

	"github.com/jinzhu/gorm"
)

type dbProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &dbProductRepository{
		DB: db,
	}
}

func (r *dbProductRepository) CreateNewProduct(p *models.Product) (*models.Product, error) {
	p.LastUpdated = time.Now()
	if dbc := r.DB.Create(p); dbc.Error != nil {
		return nil, models.ERR_PRODUCT_CONFLICT_SKU
	}
	return p, nil
}

func (r *dbProductRepository) GetDetailProduct(id int64) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	return &product, nil
}

func (r *dbProductRepository) GetDetailProductBySKU(sku string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	return &product, nil
}

func (r *dbProductRepository) FetchProduct(page int, size int) ([]models.Product, error) {
	var products []models.Product
	offset := page * size
	if err := r.DB.Offset(offset).Limit(size).Find(&products).Error; err != nil {
		return []models.Product{}, models.ERR_PRODUCT_DB
	}
	return products, nil
}

func (r *dbProductRepository) UpdateProduct(id int64, p models.Product) (*models.Product, error) {
	var existProduct models.Product
	if err := r.DB.Where("id = ?", id).First(&existProduct).Error; err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	existProduct.Name = p.Name
	existProduct.Sku = p.Sku
	existProduct.Total = p.Total
	existProduct.LastUpdated = time.Now()
	if err := r.DB.Save(&existProduct).Error; err != nil {
		return nil, models.ERR_PRODUCT_CONFLICT_SKU
	}
	return &existProduct, nil
}

func (r *dbProductRepository) UpdateProductBySKU(sku string, p models.Product) (*models.Product, error) {
	var existProduct models.Product
	if err := r.DB.Where("sku = ?", sku).First(&existProduct).Error; err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	existProduct.Name = p.Name
	existProduct.Sku = p.Sku
	existProduct.Total = p.Total
	existProduct.LastUpdated = time.Now()
	if err := r.DB.Save(&existProduct).Error; err != nil {
		return nil, models.ERR_PRODUCT_CONFLICT_SKU
	}
	return &existProduct, nil
}

func (r *dbProductRepository) DeleteProduct(id int64) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return models.ERR_PRODUCT_NOT_FOUND
	}
	return nil
}

func (r *dbProductRepository) DeleteProductBySKU(sku string) error {
	if err := r.DB.Where("sku = ?", sku).Delete(&models.Product{}).Error; err != nil {
		return models.ERR_PRODUCT_NOT_FOUND
	}
	return nil
}
