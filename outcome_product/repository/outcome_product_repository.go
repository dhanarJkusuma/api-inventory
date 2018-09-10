package repository

import (
	"fmt"
	"inventory_app/models"
	"inventory_app/outcome_product"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func (or *outcomeProductRepository) GetSalesReport(startDate string, endDate string, page int, size int) ([]models.Sales, error) {

	queryBuyPriceProduct := `
	SELECT 
		incoming_products.product_id AS product_id,
		products.name AS product_name,
		products.sku AS product_sku,
		round((sum(incoming_products.buy_price * incoming_products.total) * 1.0) / (sum(incoming_products.total)),0) AS product_buy_price
	FROM 
		incoming_products 
	JOIN products
		ON products.id = incoming_products.product_id
	GROUP BY incoming_products.product_id 
	HAVING 
		date(incoming_products.date) >= ? AND date(incoming_products.date) <= ?`

	querySales := `WITH buy_price_products AS (`
	querySales += queryBuyPriceProduct
	querySales += `)
	SELECT 
		outcoming_products.order_id AS order_id, 
		outcoming_products.date,
		buy_price_products.product_sku,
		buy_price_products.product_name,
		outcoming_products.total,
		outcoming_products.sell_price,
		round(outcoming_products.total * outcoming_products.sell_price, 0) AS total_amount,
		buy_price_products.product_buy_price as buy_price,
		round((outcoming_products.total * outcoming_products.sell_price) - (outcoming_products.total * buy_price_products.product_buy_price), 0) AS profit
	FROM 
		outcoming_products 
	JOIN buy_price_products 
		ON outcoming_products.product_id = buy_price_products.product_id
	WHERE 
		outcoming_products.order_id IS NOT NULL AND outcoming_products.order_id <> ""
	AND 
		date(outcoming_products.date) >= ? AND date(outcoming_products.date) <= ?
	ORDER BY outcoming_products.date DESC
	LIMIT ? OFFSET ?`
	fmt.Println(size)
	offset := page * size
	records := make([]models.Sales, 0)
	rows, err := or.DB.Raw(querySales, startDate, endDate, startDate, endDate, size, offset).Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println("err")
		return nil, models.ERR_RECORD_DB
	}
	for rows.Next() {
		fmt.Println("hi")
		var sales models.Sales
		err := rows.Scan(&sales.OrderID, &sales.Date, &sales.ProductSKU, &sales.ProductName, &sales.Total, &sales.SellPrice, &sales.TotalAmount, &sales.BuyPrice, &sales.Profit)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		records = append(records, sales)
	}

	return records, nil

}
