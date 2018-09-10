package http

import (
	"encoding/csv"
	"net/http"
	"path/filepath"
	"strconv"

	"inventory_app/helper"
	"inventory_app/models"
	productModule "inventory_app/product"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProductHandler struct {
	ProductUC productModule.ProductUsecase
}

func (p *ProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Product",
	}
	return c.JSON(http.StatusOK, response)
}

func (p *ProductHandler) Store(c echo.Context) error {
	var product models.Product
	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&product); !ok {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: err.Error(),
		})
	}
	insertedProduct, err := p.ProductUC.CreateNewProduct(&product)
	if err != nil {
		switch err {
		case models.ERR_PRODUCT_CONFLICT_SKU:
			return c.JSON(http.StatusConflict, &ResponseError{
				Message: "Duplicate SKU attribute",
			})
		default:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal Server Error",
			})
		}

	}
	return c.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductHandler) Fetch(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	if page = c.QueryParam("page"); page == "" {
		page = "0"
	}
	if size = c.QueryParam("size"); size == "" {
		size = "15"
	}

	if pg, err = strconv.Atoi(page); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid pagination request value",
		})
	}

	if sz, err = strconv.Atoi(size); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid pagination request value",
		})
	}
	result, err := p.ProductUC.FetchProduct(pg, sz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{
			Message: "Internal server error",
		})
	}
	return c.JSON(http.StatusOK, result)

}

func (p *ProductHandler) GetDetail(c echo.Context) error {
	// get variable path
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}

	product, err := p.ProductUC.GetProduct(id)
	if err != nil {
		// 404
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: "Product not found",
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) GetDetailBySKU(c echo.Context) error {
	// get variable path
	sku := c.Param("sku")
	product, err := p.ProductUC.GetProductBySKU(sku)
	if err != nil {
		// 404
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: "Product not found",
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) Update(c echo.Context) error {
	var updateForm models.Product

	// get variable path
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	if ok, err := isRequestValid(&updateForm); !ok {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: err.Error(),
		})
	}

	err = c.Bind(&updateForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := p.ProductUC.UpdateProduct(id, &updateForm)
	if err != nil {
		switch err {
		case models.ERR_PRODUCT_NOT_FOUND:
			return c.JSON(http.StatusNotFound, &ResponseError{
				Message: err.Error(),
			})
		case models.ERR_PRODUCT_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, result)
}

func (p *ProductHandler) UpdateBySKU(c echo.Context) error {
	var updateForm models.Product

	// get variable path
	sku := c.Param("sku")

	err := c.Bind(&updateForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := p.ProductUC.UpdateProductBySKU(sku, &updateForm)
	if err != nil {
		switch err {
		case models.ERR_PRODUCT_NOT_FOUND:
			return c.JSON(http.StatusNotFound, &ResponseError{
				Message: err.Error(),
			})
		case models.ERR_PRODUCT_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, result)
}

func (p *ProductHandler) Delete(c echo.Context) error {
	// get variable path
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	err = p.ProductUC.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	msg := struct {
		Message string
	}{
		Message: "Product has been deleted",
	}
	return c.JSON(http.StatusOK, msg)
}

func (p *ProductHandler) DeleteBySKU(c echo.Context) error {
	// get variable path
	sku := c.Param("sku")

	err := p.ProductUC.DeleteProductBySKU(sku)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	msg := struct {
		Message string
	}{
		Message: "Product has been deleted",
	}
	return c.JSON(http.StatusOK, msg)
}

func (p *ProductHandler) MigrateDataFromCSV(c echo.Context) error {
	data := make([]models.Product, 0)
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "File csv is required",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Failed to open file csv",
		})
	}
	defer src.Close()

	if filepath.Ext(file.Filename) != ".csv" {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid file type, type `csv` type required",
		})
	}

	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Broken file detected, cannot read data from csv file",
		})
	}
	var total int64

	for i, val := range lines {
		if i == 0 {
			continue
		}
		total, err = helper.StringToInt64(val[2])
		if err != nil {
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Attribute total not support.",
			})
		}

		ip := &models.Product{
			Sku:   val[0],
			Name:  val[1],
			Total: total,
		}
		data = append(data, *ip)
	}

	err = p.ProductUC.BatchInsert(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{
			Message: "Internal server error",
		})
	}
	msg := struct {
		Message string
	}{
		Message: "Data imported",
	}
	return c.JSON(http.StatusOK, msg)

}

func isRequestValid(m *models.Product) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewProductHandler(e *echo.Echo, up productModule.ProductUsecase) {
	handler := &ProductHandler{
		ProductUC: up,
	}
	e.GET("/product/ping", handler.Ping)
	e.POST("/product", handler.Store)
	e.GET("/product", handler.Fetch)
	e.GET("/product/:id", handler.GetDetail)
	e.PUT("/product/:id", handler.Update)
	e.DELETE("/product/:id", handler.Delete)
	e.GET("/product/sku/:sku", handler.GetDetailBySKU)
	e.PUT("/product/sku/:sku", handler.UpdateBySKU)
	e.DELETE("/product/sku/:sku", handler.DeleteBySKU)
	e.POST("/product/migration", handler.MigrateDataFromCSV)

}
