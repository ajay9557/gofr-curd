package main

import (
	
	"fmt"
	ProductHandler "zopsmart/productgofr/http/product"
	ProductService "zopsmart/productgofr/services/product"
	ProductStore "zopsmart/productgofr/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app:= gofr.New()

	ProStore := ProductStore.New()
	ProService := ProductService.New(ProStore)
	handler := ProductHandler.New(ProService)

	app.Server.ValidateHeaders = false
	app.Server.HTTP.Port = 8000
	app.EnableSwaggerUI()
	app.GET("/product/{id}", handler.GetProdByIdHandler)
	app.GET("/product", handler.GetAllProductHandler)
	app.POST("/product", handler.CreateProductHandler)
	app.PUT("/product/{id}", handler.UpdateProductHandler)
	app.DELETE("product/{id}", handler.DeleteProductHandler)
	fmt.Println("Listening to port : 8000")
	app.Start()
}