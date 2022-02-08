package main

import (
	productHandler "gofr-curd/delivery/product"
	productService "gofr-curd/service/product"
	productStore "gofr-curd/store/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	store := productStore.New()
	service := productService.New(store)
	h := productHandler.New(service)

	app.GET("/products/{id}", h.GetByID)
	app.GET("/products", h.Get)
	app.POST("/products", h.Create)
	app.PUT("/products/{id}", h.Update)
	app.DELETE("/products/{id}", h.Delete)

	app.Server.HTTP.Port = 8000
	app.Start()
}
