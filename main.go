package main

import (
	"fmt"
	httpProd "productGofr/http/product"
	servProd "productGofr/services/product"
	storeProd "productGofr/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	application := gofr.New()
	store := storeProd.New()
	serv := servProd.New(store)
	hndlr := httpProd.Handler{Service: serv}

	application.GET("/products/{id}", hndlr.GetByIdHandler)
	application.GET("/products", hndlr.GetAllUsers)
	application.POST("/products", hndlr.CreateProduct)
	application.DELETE("/products/{id}", hndlr.DeleteById)
	application.PUT("/products/{id}", hndlr.UpdateById)
	application.Server.HTTP.Port = 5000
	application.Server.ValidateHeaders = false
	fmt.Println("Listening to Port 5000")
	application.Start()
}
