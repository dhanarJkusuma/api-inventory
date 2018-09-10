package usecase

import (
	"inventory_app/models"
	"inventory_app/product"
)

type productUsecase struct {
	Repo product.ProductRepository
}

func NewProductUsecase(repo product.ProductRepository) product.ProductUsecase {
	return &productUsecase{
		Repo: repo,
	}
}

func (u *productUsecase) CreateNewProduct(p *models.Product) (*models.Product, error) {
	insertedProduct, err := u.Repo.CreateNewProduct(p)
	if err != nil {
		return nil, err
	}
	return insertedProduct, nil
}
func (u *productUsecase) FetchProduct(page int, size int) ([]models.Product, error) {
	result, err := u.Repo.FetchProduct(page, size)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecase) GetProduct(id int64) (*models.Product, error) {
	result, err := u.Repo.GetDetailProduct(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecase) GetProductBySKU(sku string) (*models.Product, error) {
	result, err := u.Repo.GetDetailProductBySKU(sku)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecase) UpdateProduct(id int64, p *models.Product) (*models.Product, error) {
	result, err := u.Repo.UpdateProduct(id, *p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecase) UpdateProductBySKU(sku string, p *models.Product) (*models.Product, error) {
	result, err := u.Repo.UpdateProductBySKU(sku, *p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *productUsecase) DeleteProduct(id int64) error {
	err := u.Repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) DeleteProductBySKU(sku string) error {
	err := u.Repo.DeleteProductBySKU(sku)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) BatchInsert(ps []models.Product) error {
	for _, val := range ps {
		err := u.CreateNewProductSilent(&val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *productUsecase) CreateNewProductSilent(p *models.Product) error {
	_, err := u.Repo.CreateNewProduct(p)
	if err != nil {
		return err
	}
	return nil
}
