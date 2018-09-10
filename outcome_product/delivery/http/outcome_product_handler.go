package http

import (
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"strconv"

	"inventory_app/helper"
	"inventory_app/models"
	outcomeProductModule "inventory_app/outcome_product"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type OutcomeProductHandler struct {
	OutProductUC outcomeProductModule.OutcomeProductUsecase
}

func (o *OutcomeProductHandler) Store(c echo.Context) error {
	var createForm models.OutcomingProduct

	err := c.Bind(&createForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	record, err := o.OutProductUC.CreateNewOutcomeProduct(&createForm)
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

func (o *OutcomeProductHandler) Fetch(c echo.Context) error {
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
	result, err := o.OutProductUC.FetchOutcomeProduct(from, pg, sz)
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

func (o *OutcomeProductHandler) GetDetail(c echo.Context) error {
	var id int64
	var err error

	outProductID := c.Param("id")
	if id, err = strconv.ParseInt(outProductID, 10, 64); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	result, err := o.OutProductUC.GetDetailOutcomeProduct(id)
	if err != nil {
		// 404
		return c.JSON(http.StatusNotFound, &ResponseError{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func (o *OutcomeProductHandler) Update(c echo.Context) error {
	var updateForm models.OutcomingProduct
	outProduct := c.Param("id")
	id, err := strconv.ParseInt(outProduct, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	err = c.Bind(&updateForm)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := o.OutProductUC.UpdateOutcomeProduct(id, &updateForm)
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

func (o *OutcomeProductHandler) Delete(c echo.Context) error {
	incProduct := c.Param("id")
	id, err := strconv.ParseInt(incProduct, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{
			Message: "Invalid data type path id",
		})
	}
	err = o.OutProductUC.DeleteOutcomeProduct(id)
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

func (o *OutcomeProductHandler) GetSalesReport(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	start := c.QueryParam("start")
	end := c.QueryParam("end")
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
	results, err := o.OutProductUC.GetSalesReport(start, end, pg, sz)
	if err != nil {
		switch err {
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on query start or end, `yyyy-MM-dd` required",
			})
		case models.ERR_RECORD_DB:
			return c.JSON(http.StatusInternalServerError, &ResponseError{
				Message: "Internal server error",
			})
		}
	}
	return c.JSON(http.StatusOK, results)
}

func (o *OutcomeProductHandler) ExportSalesReport(c echo.Context) error {
	var page, size string
	var pg, sz int
	var err error

	start := c.QueryParam("start")
	end := c.QueryParam("end")
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
	results, err := o.OutProductUC.GetSalesReport(start, end, pg, sz)
	if err != nil {
		switch err {
		case models.ERR_DATE_PARSING:
			return c.JSON(http.StatusBadRequest, &ResponseError{
				Message: "Error parsing date on query start or end, `yyyy-MM-dd` required",
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
	wr.Write([]string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})
	for _, val := range results {
		wr.Write([]string{val.OrderID, helper.TimeToStringFormat(val.Date, "2006-01-02 15:04:05"), val.ProductSKU, val.ProductName, helper.IntToString(val.Total), helper.FloatToString(val.SellPrice), helper.FloatToString(val.TotalAmount), helper.FloatToString(val.BuyPrice), helper.FloatToString(val.Profit)})
	}
	wr.Flush()
	return c.Attachment(f.Name(), "laporan_penjualan.csv")
}

func (o *OutcomeProductHandler) Ping(c echo.Context) error {
	response := struct {
		Message string
	}{
		Message: "Hello From Outcome Product",
	}
	return c.JSON(http.StatusOK, response)
}

func NewOutcomeProductHandler(e *echo.Echo, uop outcomeProductModule.OutcomeProductUsecase) {
	handler := &OutcomeProductHandler{
		OutProductUC: uop,
	}

	e.GET("/outcome-product/ping", handler.Ping)
	e.POST("/outcome-product", handler.Store)
	e.GET("/outcome-product", handler.Fetch)
	e.GET("/outcome-product/:id", handler.GetDetail)
	e.PUT("/outcome-product/:id", handler.Update)
	e.DELETE("/outcome-product/:id", handler.Delete)
	e.GET("/outcome-product/sales", handler.GetSalesReport)
	e.GET("/outcome-product/sales/export", handler.ExportSalesReport)
}
