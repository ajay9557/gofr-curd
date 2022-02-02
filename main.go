package main

import (
	productHandler "gofrPractice/delivery/product"
	productService "gofrPractice/service/product"
	productStore "gofrPractice/store/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	s := productStore.New()
	l := productService.New(s)
	h := productHandler.New(l)

	app.GET("/products/{id}", h.GetById)

	app.Server.HTTP.Port = 8000
	app.Start()
}
