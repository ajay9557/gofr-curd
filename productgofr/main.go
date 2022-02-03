package main

import (
	// "database/sql"
	"fmt"

	//	"developer.zopsmart.com/go/gofr/examples/sample-api/handler"
	ProductHandler "zopsmart/productgofr/handler/product"
	ProductService "zopsmart/productgofr/services/product"
	ProductStore "zopsmart/productgofr/stores/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
/*	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/user")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection establishment error")
	}
*/
	ProStore := ProductStore.New()
	ProService := ProductService.New(ProStore)
	handler := ProductHandler.New(ProService)

	app := gofr.New()
	app.Server.ValidateHeaders = false
	app.Server.HTTP.Port = 8000
	app.EnableSwaggerUI()
	app.GET("/product/{id}", handler.GetProdByIdHandler)
	fmt.Println("Listening to port : 8000")
	app.Start()
}
