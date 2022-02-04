package product

import (
	"fmt"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
)

type handler struct {
	service1 service.ProductService
}

func New(s service.ProductService) handler {
	return handler{
		service1: s,
	}
}

type response struct {
	Products []models.Product
}

func (h handler) GetProductById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return id, errors.InvalidParam{Param: []string{"id"}}
	}

	res, err := h.service1.GetProductById(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}
	fmt.Println(res)
	return res, nil
}

func (h handler) GetAllProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service1.GetAllProducts(ctx)

	if err != nil {
		return nil, errors.Error("internal error")
	}
	return resp, nil
}

func (h handler) UpdateProduct(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var prod models.Product

	if err = ctx.Bind(&prod); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	prod.Id = id

	resp, err := h.service1.UpdateProduct(ctx, prod)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) CreateProduct(ctx *gofr.Context) (interface{}, error) {
	var prod models.Product

	err := ctx.Bind(&prod)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	_, err = h.service1.CreateProduct(ctx, prod)
	if err != nil {
		return nil, errors.Error("internal errror")
	}
	return prod, nil
}

func (h handler) DeleteProduct(ctx *gofr.Context) (interface{}, error) {
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

	return "product deleted successfuly.", nil
}
