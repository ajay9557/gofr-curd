package main

import (
	httplayer "gofr-curd/http/products"
	servicelayer "gofr-curd/services/products"
	"gofr-curd/stores/products"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	store := products.New()
	service := servicelayer.New(store)
	ht := httplayer.Handler{Service: service}

	app.GET("/product/{id}", ht.GetByID)
	app.GET("/product", ht.GetAllProducts)
	app.POST("/product", ht.Insert)
	app.PUT("/product", ht.Update)
	app.DELETE("/product/{id}", ht.DeleteByID)
	app.Server.HTTP.Port = 9092
	app.Start()
}
