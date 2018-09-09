package main

import (
	"fmt"
	"inventory_app/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"

	incPrHttp "inventory_app/income_product/delivery/http"
	oucPrHttp "inventory_app/outcome_product/delivery/http"
	prHttp "inventory_app/product/delivery/http"
	trxHttp "inventory_app/transaction/delivery/http"
)

func init() {

}

func main() {
	fmt.Println("Connecting to the DB")
	// connect db
	db, err := gorm.Open("sqlite3", "./db/app.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// migration
	MigrateDatabase(db)

	// init echo
	e := echo.New()

	// handle CORS

	middlw := middleware.InitMiddleware()
	e.Use(middlw.CORS)

	// integrating module
	incPrHttp.NewIncomeProductHandler(e)
	oucPrHttp.NewOutcomeProductHandler(e)
	prHttp.NewProductHandler(e)
	trxHttp.NewTransactionHandler(e)

	e.Logger.Fatal(e.Start(":8080"))

}
