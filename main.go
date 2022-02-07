package main

import (
	httplayer "gofr-curd/Http/product"
	servicelayer "gofr-curd/service/product"
	"gofr-curd/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	store := product.New()
	Service := servicelayer.New(&store)
	ht := httplayer.Handler{Service}

	app.GET("/product/{id}", ht.GetById)
	app.GET("/product", ht.GetAllProducts)
	app.POST("/product", ht.Insert)
	app.PUT("/product", ht.Update)
	app.DELETE("/product/{id}", ht.DeleteById)
	app.Server.HTTP.Port = 9092
	app.Start()

}
