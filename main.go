package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	handler "github.com/himanshu-kumar-zs/gofr-curd/handler"
	service "github.com/himanshu-kumar-zs/gofr-curd/services/product"
	store "github.com/himanshu-kumar-zs/gofr-curd/store/product"
)

func main() {
	app := gofr.New()
	store := store.New()
	serv := service.New(store)
	handler := handler.New(serv)
	app.GET("/product/{id}", handler.GetByID)
	app.GET("/products", handler.GetAll)
	app.PUT("/product/{id}", handler.Update)
	app.POST("/product", handler.Create)
	app.DELETE("/product/{id}", handler.Delete)
	app.Server.HTTP.Port = 3000
	app.Server.ValidateHeaders = false
	app.Start()
}
