package usecase

import (
	"inventory_app/income_product"
	"inventory_app/models"
	"inventory_app/product"
	"time"
)

type incomeProductUsecase struct {
	ProductRepo       product.ProductRepository
	IncomeProductRepo income_product.IncomeProductRepository
}

func NewIncomeProductUsecase(prRepo product.ProductRepository, incPrRepo income_product.IncomeProductRepository) income_product.IncomeProductUsecase {
	return &incomeProductUsecase{
		ProductRepo:       prRepo,
		IncomeProductRepo: incPrRepo,
	}
}

func (u *incomeProductUsecase) CreateNewIncomeProduct(ip *models.IncomingProduct) (*models.IncomingProduct, error) {
	p, err := u.ProductRepo.GetDetailProductBySKU(ip.ProductSKU)
	if err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	ip.ProductID = p.ID
	ip.ProductSKU = p.Sku
	ip.Product = *p

	// check date
	parsedDate, err := time.Parse("2006-01-02 15:04:05", ip.DateFormatted)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	ip.Date = parsedDate

	incomeProduct, err := u.IncomeProductRepo.CreateNewIncomeProduct(ip)
	if err != nil {
		return nil, models.ERR_RECORD_DB
	}
	// update stock
	p.Total += ip.Total
	u.ProductRepo.UpdateProduct(p.ID, *p)
	ip.Product = *p
	return incomeProduct, nil
}

func (u *incomeProductUsecase) FetchIncomeProduct(from string, page int, size int) ([]models.IncomingProduct, error) {
	if from == "" {
		from = "1980-01-01"
	}
	fromDate, err := time.Parse("2006-01-02", from)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	result, err := u.IncomeProductRepo.FetchIncomeProduct(fromDate, page, size)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *incomeProductUsecase) GetDetailIncomeProduct(id int64) (*models.IncomingProduct, error) {
	result, err := u.IncomeProductRepo.GetDetailIncomeProduct(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *incomeProductUsecase) GetDetailIncomeProductByNoReceipt(no string) (*models.IncomingProduct, error) {
	result, err := u.IncomeProductRepo.GetDetailIncomeProductByNoReceipt(no)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *incomeProductUsecase) UpdateIncomeProduct(id int64, ip *models.IncomingProduct) (*models.IncomingProduct, error) {
	p, err := u.ProductRepo.GetDetailProductBySKU(ip.ProductSKU)
	if err != nil {
		return nil, models.ERR_PRODUCT_NOT_FOUND
	}
	ip.ProductID = p.ID
	ip.ProductSKU = p.Sku
	ip.Product = *p

	// check date
	parsedDate, err := time.Parse("2006-01-02 15:04:05", ip.DateFormatted)
	if err != nil {
		return nil, models.ERR_DATE_PARSING
	}
	ip.Date = parsedDate

	incomeProduct, oldData, err := u.IncomeProductRepo.UpdateIncomeProduct(id, *ip)
	if err != nil {
		return nil, models.ERR_RECORD_DB
	}
	// check stock changing
	if oldData.ProductID == p.ID {
		delta := ip.Total - oldData.Total
		p.Total += delta
		// update stock
		u.ProductRepo.UpdateProduct(p.ID, *p)
	} else {
		// update old stock
		oldData.Product.Total -= oldData.Total
		u.ProductRepo.UpdateProduct(oldData.Product.ID, oldData.Product)

		// update new stock
		p.Total += ip.Total
		u.ProductRepo.UpdateProduct(p.ID, *p)
	}
	incomeProduct.Product = *p
	return incomeProduct, nil
}

func (u *incomeProductUsecase) DeleteIncomeProduct(id int64) error {
	oldData, err := u.IncomeProductRepo.GetDetailIncomeProduct(id)
	if err != nil {
		return err
	}
	err = u.IncomeProductRepo.DeleteIncomeProduct(id)
	if err != nil {
		return err
	}
	// update old stock
	oldData.Product.Total -= oldData.Total
	u.ProductRepo.UpdateProduct(oldData.Product.ID, oldData.Product)

	return nil
}
