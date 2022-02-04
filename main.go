package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	productHandler "github.com/tejas/gofr-crud/http/product"
	productService "github.com/tejas/gofr-crud/service/product"
	productStore "github.com/tejas/gofr-crud/store/product"
)

func main() {
	app := gofr.New()

	prodStore := productStore.New()
	prodService := productService.New(prodStore)
	prodHandler := productHandler.New(prodService)

	app.GET("/product/{id}", prodHandler.GetProductById)
	app.GET("/products", prodHandler.GetAllProducts)
	app.PUT("/product/{id}", prodHandler.UpdateProduct)
	app.DELETE("/product/{id}", prodHandler.DeleteProduct)
	app.POST("/product", prodHandler.CreateProduct)

	app.Server.ValidateHeaders = false

	app.Server.HTTP.Port = 8000
	app.Start()

}
