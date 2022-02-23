package main

import (
	"fmt"
	Handler "gofr-curd/http/product"
	serviceHandler "gofr-curd/service/product"
	storeHandler "gofr-curd/store/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	storeH := storeHandler.New()
	serviceH := serviceHandler.New(storeH)
	handler := Handler.New(serviceH)

	app.GET("/product/{id}", handler.GetByID)
	app.GET("/product", handler.GetAllProductDetails)
	app.POST("/product", handler.InsertProduct)
	app.PUT("/product", handler.UpdateProductByID)
	app.DELETE("/product/{id}", handler.DeleteByProductID)
	app.Server.HTTP.Port = 8000

	fmt.Println("Listening at port 8000")

	app.Start()
}
