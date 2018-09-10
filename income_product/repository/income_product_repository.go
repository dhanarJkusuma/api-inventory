package repository

import (
	"fmt"
	"inventory_app/income_product"
	"inventory_app/models"
	"time"

	"github.com/jinzhu/gorm"
)

type incomeProductRepository struct {
	DB *gorm.DB
}

func NewIncomeProductRepository(db *gorm.DB) income_product.IncomeProductRepository {
	return &incomeProductRepository{
		DB: db,
	}
}

func (i *incomeProductRepository) CreateNewIncomeProduct(ip *models.IncomingProduct) (*models.IncomingProduct, error) {
	ip.LastUpdated = time.Now()
	ip.InsertedAt = time.Now()
	ip.TotalAmount = float64(ip.TotalOrder) * float64(ip.BuyPrice)
	if dbc := i.DB.Create(ip); dbc.Error != nil {
		return nil, models.ERR_RECORD_DB
	}
	i.DB.Model(&ip).Related(&ip.Product)
	return ip, nil
}

func (i *incomeProductRepository) FetchIncomeProduct(from time.Time, page int, size int) ([]models.IncomingProduct, error) {
	var incomeProducts []models.IncomingProduct
	offset := page * size
	if err := i.DB.Where("date >= ? ", from).Order("date desc").Offset(offset).Limit(size).Find(&incomeProducts).Error; err != nil {
		fmt.Println("error")
		return []models.IncomingProduct{}, models.ERR_RECORD_DB
	}
	for index := range incomeProducts {
		i.DB.Model(incomeProducts[index]).Related(&incomeProducts[index].Product)
		incomeProducts[index].ProductSKU = incomeProducts[index].Product.Sku
	}
	return incomeProducts, nil
}

func (i *incomeProductRepository) GetDetailIncomeProduct(id int64) (*models.IncomingProduct, error) {
	var incomeProduct models.IncomingProduct
	if err := i.DB.Where("id = ?", id).First(&incomeProduct).Error; err != nil {
		return nil, models.ERR_RECORD_NOT_FOUND
	}
	i.DB.Model(&incomeProduct).Related(&incomeProduct.Product)
	return &incomeProduct, nil
}

func (i *incomeProductRepository) GetDetailIncomeProductByNoReceipt(no string) (*models.IncomingProduct, error) {
	var incomeProduct models.IncomingProduct
	if err := i.DB.Where("number_receipt = ?", no).First(&incomeProduct).Error; err != nil {
		return nil, models.ERR_RECORD_NOT_FOUND
	}
	i.DB.Model(&incomeProduct).Related(&incomeProduct.Product)
	return &incomeProduct, nil
}

func (i *incomeProductRepository) UpdateIncomeProduct(id int64, p models.IncomingProduct) (*models.IncomingProduct, models.IncomingProduct, error) {
	var existIncProd models.IncomingProduct
	if err := i.DB.Where("id = ?", id).First(&existIncProd).Error; err != nil {
		return nil, models.IncomingProduct{}, models.ERR_RECORD_NOT_FOUND
	}
	i.DB.Model(&existIncProd).Related(&existIncProd.Product)
	oldData := existIncProd
	existIncProd.ProductID = p.ProductID
	existIncProd.ProductSKU = p.ProductSKU
	existIncProd.Product = p.Product
	existIncProd.TotalOrder = p.TotalOrder
	existIncProd.Total = p.Total
	existIncProd.Note = p.Note
	existIncProd.BuyPrice = p.BuyPrice
	existIncProd.TotalAmount = float64(p.TotalOrder) * float64(p.BuyPrice)
	existIncProd.NumberReceipt = p.NumberReceipt
	existIncProd.LastUpdated = time.Now()
	if err := i.DB.Save(&existIncProd).Error; err != nil {
		return nil, models.IncomingProduct{}, models.ERR_RECORD_DB
	}
	return &existIncProd, oldData, nil
}

func (i *incomeProductRepository) DeleteIncomeProduct(id int64) error {
	if err := i.DB.Where("id = ?", id).Delete(&models.IncomingProduct{}).Error; err != nil {
		return models.ERR_RECORD_DB
	}
	return nil
}

func (i *incomeProductRepository) GetSummaryProductValue(date time.Time, page int, size int) ([]models.ProductValue, error) {
	var results []models.ProductValue

	querySelect := "incoming_products.product_id AS product_id, "
	querySelect += "products.sku AS product_sku, "
	querySelect += "products.name AS product_name, "
	querySelect += "sum(incoming_products.total) as total, "
	querySelect += "round(sum(incoming_products.buy_price * incoming_products.total), 0) AS total_buy_price, "
	querySelect += "round((sum(incoming_products.buy_price * incoming_products.total) * 1.0) / (sum(incoming_products.total)), 0) AS avg_buy_price,"
	querySelect += "round((sum(incoming_products.total) * 1.0) * (sum(incoming_products.buy_price * incoming_products.total) * 1.0) / (sum(incoming_products.total)), 0) AS total_amount, "
	querySelect += "incoming_products.date as date"

	queryJoin := "join products "
	queryJoin += "on products.id = incoming_products.product_id"

	groupingBy := "incoming_products.product_id, products.sku"

	havingAttribute := "date(incoming_products.date) >= ?"
	havingValue := date

	offset := page * size

	err := i.DB.Table("incoming_products").Select(querySelect).Joins(queryJoin).Group(groupingBy).Having(havingAttribute, havingValue).Offset(offset).Limit(size).Scan(&results).Error
	if err != nil {
		fmt.Println(err)
		return []models.ProductValue{}, models.ERR_RECORD_DB
	}

	return results, nil
}
