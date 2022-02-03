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
	application.Server.HTTP.Port = 5000
	application.Server.ValidateHeaders = false
	fmt.Println("Listening to Port 5000")
	application.Start()

}
