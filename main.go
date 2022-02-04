package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/handler"
	"github.com/himanshu-kumar-zs/gofr-curd/store/product"
)

func main() {
	app := gofr.New()
	store := product.New()
	handler := handler.New(store)
	app.GET("/product/{id}", handler.GetByID)
	app.Server.HTTP.Port = 3000
	app.Server.ValidateHeaders = false
	app.Start()
}
