package product

import (
	"fmt"
	"gofr-curd/models"
	"gofr-curd/services"
	"net/http"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Service services.Iservice
}

func (h Handler) GetProductByIDHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")

	p, err := h.Service.GetProductByID(ctx, param)

	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Data:       &p,
		Message:    "Successfully Product Fetched",
		StatusCode: http.StatusOK,
	}

	return responseObj, nil
}

func (h Handler) GetAllProductsHandler(ctx *gofr.Context) (interface{}, error) {
	var prds []*models.Product

	products, err := h.Service.GetAllProducts(ctx)

	if err != nil {
		return prds, err
	}

	prds = products
	responseObj := models.Response{
		Data:       &prds,
		Message:    "Successfully Product Fetched",
		StatusCode: http.StatusOK,
	}

	return responseObj, nil
}

func (h Handler) CreateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var prd models.Product
	err := ctx.Bind(&prd)

	if err != nil {
		fmt.Print("Anusri")
		ctx.Logger.Errorf("error in binding: %v", err)

		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	NewPrd, err := h.Service.CreateProduct(ctx, prd)

	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Data:       &NewPrd,
		Message:    "Successfully Product Created",
		StatusCode: http.StatusOK,
	}

	return responseObj, nil
}

func (h Handler) DeleteByIDHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")
	err := h.Service.DeleteByID(ctx, param)

	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Message:    "Successfully Product Deleted",
		StatusCode: http.StatusOK,
	}

	return responseObj, nil
}

func (h Handler) UpdateByIDHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")

	var prd models.Product
	err := ctx.Bind(&prd)

	if err != nil {
		fmt.Print("Anusri")
		ctx.Logger.Errorf("error in binding: %v", err)

		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	p, err := h.Service.UpdateByID(ctx, param, prd)
	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Data:       &p,
		Message:    "Successfully Product Updated",
		StatusCode: http.StatusOK,
	}

	return responseObj, nil
}
