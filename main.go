package main

import (
	productHandler "zopsmart/gofr-curd/handler/products"
	productService "zopsmart/gofr-curd/service/products"
	productStore "zopsmart/gofr-curd/store/products"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	store := productStore.New()
	service := productService.New(&store)
	handler := productHandler.New(service)

	app.GET("/product/{id}", handler.GetByID)
	app.GET("/products", handler.GetProducts)
	app.POST("/products", handler.AddProduct)
	app.DELETE("/products/{id}", handler.DeleteByID)
	app.PUT("/products/{id}", handler.UpdateByID)

	app.Server.HTTP.Port = 8000
	app.Start()
}
