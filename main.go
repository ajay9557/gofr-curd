package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	handler "github.com/himanshu-kumar-zs/gofr-curd/handler"
	service "github.com/himanshu-kumar-zs/gofr-curd/services/product"
	store "github.com/himanshu-kumar-zs/gofr-curd/store/product"
)

func main() {
	app := gofr.New()
	productStore := store.New()
	serv := service.New(productStore)
	h := handler.New(serv)
	app.GET("/product/{id}", h.GetByID)
	app.GET("/products", h.GetAll)
	app.PUT("/product/{id}", h.Update)
	app.POST("/product", h.Create)
	app.DELETE("/product/{id}", h.Delete)
	app.Server.HTTP.Port = 3000
	app.Server.ValidateHeaders = false
	app.Start()
}
