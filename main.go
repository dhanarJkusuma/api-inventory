package main

import (
	"fmt"
	"inventory_app/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"

	incPrHttp "inventory_app/income_product/delivery/http"
	iprRepo "inventory_app/income_product/repository"
	iprUsecase "inventory_app/income_product/usecase"

	oucPrHttp "inventory_app/outcome_product/delivery/http"
	oprRepo "inventory_app/outcome_product/repository"
	oprUsecase "inventory_app/outcome_product/usecase"

	prHttp "inventory_app/product/delivery/http"
	prRepo "inventory_app/product/repository"
	prUsecase "inventory_app/product/usecase"
)

func init() {

}

func main() {
	fmt.Println("[Inventory App] Try connecting to the Database")
	// connect db
	db, err := gorm.Open("sqlite3", "./db/app.db")
	if err != nil {
		panic("[Inventory App] Failed to connect Database")
	}
	defer db.Close()
	fmt.Println("[Inventory App] Connected to the Database")

	// migration
	MigrateDatabase(db)

	// init echo
	e := echo.New()

	// handle CORS

	middlw := middleware.InitMiddleware()
	e.Use(middlw.CORS)

	// connecting module product
	pr := prRepo.NewProductRepository(db)
	pu := prUsecase.NewProductUsecase(pr)
	prHttp.NewProductHandler(e, pu)

	// connecting module incoming product
	ipr := iprRepo.NewIncomeProductRepository(db)
	ipu := iprUsecase.NewIncomeProductUsecase(pr, ipr)
	incPrHttp.NewIncomeProductHandler(e, ipu)

	// connecting module outcoming product
	opr := oprRepo.NewOutcomeProductRepository(db)
	opu := oprUsecase.NewOutcomeProductUsecase(pr, opr)
	oucPrHttp.NewOutcomeProductHandler(e, opu)

	fmt.Println("[Inventory App] Application Starting")
	e.Logger.Fatal(e.Start(":8080"))

}
