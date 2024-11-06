package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"product-app/common/app"
	"product-app/common/postgresql"
	"product-app/controller"
	"product-app/persistence"
	"product-app/service"
)

func main() {
	ctx := context.Background()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := controller.NewProductController(productService)
	e := echo.New()

	productController.RegisterRoutes(e)

	err := e.Start("127.0.0.1:9650")
	if err != nil {
		fmt.Println("Error while starting server {}", err.Error())
	}

}
