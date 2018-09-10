package usecase

import (
	"inventory_app/helper"
	"inventory_app/models"
	"inventory_app/outcome_product"
	"inventory_app/product"
	"time"
)

type outcomeProductUsecase struct {
	ProductRepo        product.ProductRepository
	OutcomeProductRepo outcome_product.OutcomeProductRepository
}

func NewOutcomeProductUsecase(prRepo product.ProductRepository, outPrRepo outcome_product.OutcomeProductRepository) outcome_product.OutcomeProductUsecase {
	return &outcomeProductUsecase{
		ProductRepo:        prRepo,
		OutcomeProductRepo: outPrRepo,
	}
}

func (ou *outcomeProductUsecase) CreateNewOutcomeProduct(op *models.OutcomingProduct) (*models.OutcomingProduct, error) {
	p, err := ou.ProductRepo.GetDetailProductBySKU(op.ProductSKU)
	if err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	op.ProductID = p.ID
	op.ProductSKU = p.Sku
	op.Product = *p

	// check date
	parsedDate, err := time.Parse("2006-01-02 15:04:05", op.DateFormatted)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	op.Date = parsedDate

	outcomeProduct, err := ou.OutcomeProductRepo.CreateNewOutcomeProduct(op)
	if err != nil {
		return nil, models.ERR_RECORD_DB
	}
	// update stock
	p.Total -= op.Total
	ou.ProductRepo.UpdateProduct(p.ID, *p)
	outcomeProduct.Product = *p
	return outcomeProduct, nil
}

func (ou *outcomeProductUsecase) FetchOutcomeProduct(from string, page int, size int) ([]models.OutcomingProduct, error) {
	if from == "" {
		from = "1980-01-01"
	}
	fromDate, err := time.Parse("2006-01-02", from)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	result, err := ou.OutcomeProductRepo.FetchOutcomeProduct(fromDate, page, size)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ou *outcomeProductUsecase) GetDetailOutcomeProduct(id int64) (*models.OutcomingProduct, error) {
	result, err := ou.OutcomeProductRepo.GetDetailOutcomeProduct(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ou *outcomeProductUsecase) UpdateOutcomeProduct(id int64, op *models.OutcomingProduct) (*models.OutcomingProduct, error) {
	p, err := ou.ProductRepo.GetDetailProductBySKU(op.ProductSKU)
	if err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	op.ProductID = p.ID
	op.ProductSKU = p.Sku
	op.Product = *p

	// check date
	parsedDate, err := time.Parse("2006-01-02 15:04:05", op.DateFormatted)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	op.Date = parsedDate

	outcomeProduct, oldData, err := ou.OutcomeProductRepo.UpdateOutcomeProduct(id, *op)
	if err != nil {
		return nil, models.ERR_RECORD_DB
	}
	// check stock changing
	if oldData.ProductID == p.ID {
		delta := op.Total - oldData.Total
		p.Total -= delta
		// update stock
		ou.ProductRepo.UpdateProduct(p.ID, *p)
	} else {
		// update old stock
		oldData.Product.Total += oldData.Total
		ou.ProductRepo.UpdateProduct(oldData.Product.ID, oldData.Product)

		// update new stock
		p.Total -= op.Total
		ou.ProductRepo.UpdateProduct(p.ID, *p)
	}
	outcomeProduct.Product = *p
	return outcomeProduct, nil
}

func (ou *outcomeProductUsecase) DeleteOutcomeProduct(id int64) error {
	oldData, err := ou.OutcomeProductRepo.GetDetailOutcomeProduct(id)
	if err != nil {
		return err
	}
	err = ou.OutcomeProductRepo.DeleteOutcomeProduct(id)
	if err != nil {
		return err
	}
	// update old stock
	oldData.Product.Total += oldData.Total
	ou.ProductRepo.UpdateProduct(oldData.Product.ID, oldData.Product)

	return nil
}

func (ou *outcomeProductUsecase) GetSalesReport(startDate string, endDate string, page int, size int) ([]models.Sales, error) {
	if startDate == "" {
		startDate = "1980-01-01"
	}
	if endDate == "" {
		endDate = helper.GetCurrentDateWithFormat("2006-01-02")
	}
	err := helper.IsCorrectDateFormat("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	err = helper.IsCorrectDateFormat("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}
	results, err := ou.OutcomeProductRepo.GetSalesReport(startDate, endDate, page, size)
	if err != nil {
		return nil, err
	}
	return results, nil
}
