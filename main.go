package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	pHttp "github.com/arohanzst/testapp/http/products"
	pServices "github.com/arohanzst/testapp/services/products"
	pStore "github.com/arohanzst/testapp/stores/products"
)

func main() {
	app := gofr.New()

	myStore := pStore.New()
	myService := pServices.New(myStore)
	handler := pHttp.New(myService)

	// Specifying the different routes supported by this service
	app.GET("/product/{id}", handler.ReadByIDHandler)
	app.GET("/product", handler.ReadHandler)
	app.POST("/product", handler.CreateHandler)
	app.PUT("/product/{id}", handler.UpdateHandler)
	app.DELETE("/product/{id}", handler.DeleteHandler)

	// starting the server on a custom port
	app.Server.HTTP.Port = 8080
	app.Server.ValidateHeaders = false
	app.Start()
}
