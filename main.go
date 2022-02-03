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
	ht := httplayer.Handler{service}

	app.GET("/product/{id}", ht.GetById)
	app.Server.HTTP.Port = 9092
	app.Start()

}
