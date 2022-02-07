package product

import (
	"net/http"
	"product/models"
	"product/services"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type HttpService struct {
	Service services.Service
}

func New(service services.Service) HttpService {
	return HttpService{Service: service}
}

func (service *HttpService) GetByIdHandler(ctx *gofr.Context) (interface{}, error) {
	productId := ctx.PathParam("id")
	var product models.Product

	id, _ := strconv.Atoi(productId)

	product, err := service.Service.GetProductById(ctx, id)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     productId,
		}
	}

	responseObj := models.Response{
		Data:       product,
		Message:    "Product Found",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}

func (service *HttpService) GetAllProductHandler(ctx *gofr.Context) (interface{}, error) {

	productList, _ := service.Service.GetAllProduct(ctx)

	responseObj := models.Response{
		Data:       productList,
		Message:    "Products Found",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}

func (service *HttpService) AddProductHandler(ctx *gofr.Context) (interface{}, error) {

	var product models.Product

	_ = ctx.Bind(&product)

	err := service.Service.AddProduct(ctx, product)

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "FAILED TO ADD PRODUCT", ID: ""}
	}

	responseObj := models.Response{
		Data:       "Product Added",
		Message:    "Saved",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}

func (service *HttpService) UpdateProductHandler(ctx *gofr.Context) (interface{}, error) {

	var product models.Product

	_ = ctx.Bind(&product)

	err := service.Service.UpdateProduct(ctx, product)

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "FAILED TO UPDATE PRODUCT", ID: ""}
	}

	responseObj := models.Response{
		Data:       "Product Updated",
		Message:    "Successfull",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}

func (service *HttpService) DeleteProductHandler(ctx *gofr.Context) (interface{}, error) {

	productId := ctx.PathParam("id")

	id, err := strconv.Atoi(productId)
	if err != nil || id < 0 {
		return nil, errors.EntityNotFound{Entity: "INVALID INPUTS", ID: ""}
	}

	_ = service.Service.DeleteProduct(ctx, id)

	responseObj := models.Response{
		Data:       "Product Deleted",
		Message:    "Successfull",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}
