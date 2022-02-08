package product

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
)

type Handler struct {
	service1 service.ProductService
}

func New(s service.ProductService) Handler {
	return Handler{service1: s}
}

func (h Handler) GetProductByID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return id, errors.InvalidParam{Param: []string{"id"}}
	}

	res, err := h.service1.GetProductByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}
	var resp models.Response
	resp = models.Response{
		Product:    res,
		Message:    "product data fetching successful",
		StatusCode: 200,
	}

	return resp, nil
}

func (h Handler) GetAllProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service1.GetAllProducts(ctx)

	if err != nil {
		return nil, errors.Error("internal database error")
	}

	var res models.Response
	res = models.Response{
		Product:    resp,
		Message:    "product data fetching successful",
		StatusCode: 200,
	}

	return res, nil
}

func (h Handler) UpdateProduct(ctx *gofr.Context) (interface{}, error) {
	var prod models.Product

	err1 := ctx.Bind(&prod)
	if err1 != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	_, err1 = h.service1.UpdateProduct(ctx, prod)

	if err1 != nil {
		return nil, errors.Error("internal error")
	}
	var resp models.Response
	resp = models.Response{
		Product:    prod,
		Message:    "successfully updated product data",
		StatusCode: 200,
	}

	return resp, nil
}

func (h Handler) CreateProduct(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	err := ctx.Bind(&product)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	_, err = h.service1.CreateProduct(ctx, product)

	if err != nil {
		return nil, errors.Error("internal errror")
	}
	var resp models.Response
	resp = models.Response{
		Product:    product,
		Message:    "product creation successful",
		StatusCode: 200,
	}

	return resp, nil
}

func (h Handler) DeleteProduct(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.service1.DeleteProduct(ctx, id); err != nil {
		return nil, err
	}
	var res models.Response
	res = models.Response{
		Product:    nil,
		Message:    "deleted product data successfully",
		StatusCode: 200,
	}

	return res, nil
}
