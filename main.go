package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	service "github.com/shaurya-zopsmart/Gopr-devlopment/Services/user"
	store "github.com/shaurya-zopsmart/Gopr-devlopment/Stores/PRODUCT"
	http "github.com/shaurya-zopsmart/Gopr-devlopment/http/user"
)

func main() {
	app := gofr.New()

	st := store.New()
	se := service.New(st)
	handler := http.New(se)
	app.GET("/dev/{id}", handler.GetByID)
	app.POST("/dev", handler.Create)
	app.DELETE("/dev/{id}", handler.Delete)
	app.PUT("/dev/{id}", handler.Update)
	app.GET("/dev", handler.GetAll)

	app.Server.ValidateHeaders = false

	app.Server.HTTP.Port = 9092

	app.Start()
}
