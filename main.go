package main

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	productHandler "github.com/ridhdhish-desai-zs/product-gofr/http/products"
	productService "github.com/ridhdhish-desai-zs/product-gofr/service/products"
	"github.com/ridhdhish-desai-zs/product-gofr/store/products"
)

func main() {
	app := gofr.New()

	st := products.New()
	srv := productService.New(st)
	h := productHandler.New(srv)

	app.GET("/products/{id}", h.GetByIdHandler)
	app.GET("/products", h.GetHandler)

	app.Server.HTTP.Port = 5000
	app.Server.ValidateHeaders = false

	fmt.Println("Listening to PORT 5000")
	app.Start()
}
