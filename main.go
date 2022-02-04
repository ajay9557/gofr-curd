package main

import (
	"fmt"
	httpProd "gofr-curd/http/product"
	servProd "gofr-curd/services/product"
	storeProd "gofr-curd/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	application := gofr.New()
	store := storeProd.New()
	serv := servProd.New(store)
	hndlr := httpProd.Handler{Service: serv}

	application.GET("/products/{id}", hndlr.GetProductByIdHandler)
	application.GET("/products", hndlr.GetAllProductsHandler)
	application.POST("/products", hndlr.CreateProductHandler)
	application.DELETE("/products/{id}", hndlr.DeleteByIdHandler)
	application.PUT("/products/{id}", hndlr.UpdateByIdHandler)
	application.Server.HTTP.Port = 5000
	application.Server.ValidateHeaders = false
	fmt.Println("Listening to Port 5000")
	application.Start()
}
