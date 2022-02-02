package main

import (
	httpProd "product/http/product"
	servProd "product/services/product"
	storeProd "product/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	application := gofr.New()

	store := storeProd.New()
	serv := servProd.New(store)
	handler := httpProd.HttpService{Service: serv}

	application.Server.ValidateHeaders = false

	application.GET("/product/{id}", handler.GetByIdHandler)

	application.Start()
}
