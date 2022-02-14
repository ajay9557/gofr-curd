package main

import (
	httplayer "github.com/Training/gofr-curd/http/products"
	servicelayer "github.com/Training/gofr-curd/services/products"
	"github.com/Training/gofr-curd/stores/products"

	"developer.zopsmart.com/go/gofr/examples/sample-api/handler"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	store := products.New()
	service := servicelayer.New(store)
	ht := httplayer.Handler{Service: service}

	app.EnableSwaggerUI()

	app.GET("/product/{id}", ht.GetByID)
	app.GET("/product", ht.GetAllProducts)
	app.POST("/product", ht.Insert)
	app.PUT("/product", ht.Update)
	app.DELETE("/product/{id}", ht.DeleteByID)
	app.GET("/json", handler.JSONHandler)
	app.Server.HTTP.Port = 9092
	app.Start()
}
