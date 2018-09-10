package http

import (
	"encoding/csv"
	"inventory_app/helper"
	incomeProductModule "inventory_app/income_product"
	"inventory_app/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type IncomeProductHandler struct {
	IncProductUC incomeProductModule.IncomeProductUsecase
}

func (i *IncomeProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Income Product",
	}
	return c.JSON(http.StatusOK, response)
}

func (i *IncomeProductHandler) Store(c echo.Context) error {
	var createForm models.IncomingProduct

	err := c.Bind(&createForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	record, err := i.IncProductUC.CreateNewIncomeProduct(&createForm)
	if err != nil {
		switch err {
		case models.ERR_PRODUCT_NOT_FOUND:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "ProductSKU attribute doesn't exist",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on attribute dateFormatted, `yyyy/MM/dd HH:mm` required",
			})
		}
	}
	return c.JSON(http.StatusCreated, record)
}

func (i *IncomeProductHandler) Fetch(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	from := c.QueryParam("from")
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
	result, err := i.IncProductUC.FetchIncomeProduct(from, pg, sz)
	if err != nil {
		switch err {
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Invalid parameter `from` format, need yyyy-mm-dd date formatted",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, result)
}

func (i *IncomeProductHandler) GetDetail(c echo.Context) error {
	var id int64
	var err error

	incProduct := c.Param("id")
	if id, err = strconv.ParseInt(incProduct, 10, 64); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	result, err := i.IncProductUC.GetDetailIncomeProduct(id)
	if err != nil {
		// 404
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func (i *IncomeProductHandler) GetDetailByNoReceipt(c echo.Context) error {
	var err error

	noReceipt := c.Param("no")
	result, err := i.IncProductUC.GetDetailIncomeProductByNoReceipt(noReceipt)
	if err != nil {
		// 404
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func (i *IncomeProductHandler) Update(c echo.Context) error {
	var updateForm models.IncomingProduct
	incProduct := c.Param("id")
	id, err := strconv.ParseInt(incProduct, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	err = c.Bind(&updateForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := i.IncProductUC.UpdateIncomeProduct(id, &updateForm)
	if err != nil {
		switch err {
		case models.ERR_PRODUCT_NOT_FOUND:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "ProductSKU attribute doesn't exist",
			})
		case models.ERR_RECORD_NOT_FOUND:
			return c.JSON(http.StatusNotFound, &ResponseError{
				Message: err.Error(),
			})
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on attribute dateFormatted, `yyyy/MM/dd HH:mm` required",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, result)
}

func (i *IncomeProductHandler) Delete(c echo.Context) error {
	incProduct := c.Param("id")
	id, err := strconv.ParseInt(incProduct, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	err = i.IncProductUC.DeleteIncomeProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	msg := struct {
		Message string
	}{
		Message: "Record has been deleted",
	}
	return c.JSON(http.StatusOK, msg)
}

func (i *IncomeProductHandler) GetSummaryProductValue(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	from := c.QueryParam("from")
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

	results, err := i.IncProductUC.GetSummaryProductValue(from, pg, sz)
	if err != nil {
		switch err {
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on attribute dateFormatted, `yyyy/MM/dd HH:mm` required",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, results)
}

func (i *IncomeProductHandler) ExportSummaryProductValue(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	from := c.QueryParam("from")
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

	results, err := i.IncProductUC.GetSummaryProductValue(from, pg, sz)
	if err != nil {
		switch err {
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on attribute dateFormatted, `yyyy/MM/dd HH:mm` required",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}

	f, err := ioutil.TempFile("/tmp", "tmp")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{
			Message: "Internal server error",
		})
	}
	wr := csv.NewWriter(f)

	// writing header
	wr.Write([]string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})
	for _, val := range results {
		wr.Write([]string{val.ProductSKU, val.ProductName, helper.IntToString(val.Total), helper.FloatToString(val.AvgBuyPrice), helper.FloatToString(val.TotalAmount)})
	}
	wr.Flush()

	return c.Attachment(f.Name(), "laporan_nilai_barang.csv")
}

func NewIncomeProductHandler(e *echo.Echo, uip incomeProductModule.IncomeProductUsecase) {
	handler := &IncomeProductHandler{
		IncProductUC: uip,
	}
	e.GET("/income-product/ping", handler.Ping)
	e.POST("/income-product", handler.Store)
	e.GET("/income-product", handler.Fetch)
	e.GET("/income-product/:id", handler.GetDetail)
	e.PUT("/income-product/:id", handler.Update)
	e.DELETE("/income-product/:id", handler.Delete)
	e.GET("/income-product/product-value", handler.GetSummaryProductValue)
	e.GET("/income-product/product-value/export", handler.ExportSummaryProductValue)
}
