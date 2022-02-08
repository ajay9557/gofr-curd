package products

import (
	"fmt"
	"gofr-curd/models"
	"gofr-curd/services"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Service services.Service
}

func (h Handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	var result models.Response

	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.Service.GetByUserID(ctx, id)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Product",
			ID:     i,
		}
	}

	result = models.Response{
		Product:    resp,
		Message:    "Retreived Product Successfully",
		StatusCode: 200,
	}

	return result, nil
}

func (h Handler) DeleteByID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	var res models.Response

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = h.Service.DeleteByProductID(ctx, id)

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
	}

	res = models.Response{
		Product:    nil,
		Message:    "Deleted Product Successfully",
		StatusCode: 200,
	}

	return res, nil
}

func (h Handler) Update(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	var res models.Response

	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	err := h.Service.UpdateByProductID(ctx, product)

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	res = models.Response{
		Product:    product,
		Message:    "Updated Product Successfully",
		StatusCode: 200,
	}

	return res, nil
}

func (h Handler) Insert(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	var result models.Response

	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if product.ID == 0 {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.Service.InsertProduct(ctx, product)

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	result = models.Response{
		Product:    resp,
		Message:    "Inserted Product Successfully",
		StatusCode: 200,
	}

	return result, nil
}

func (h Handler) GetAllProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Service.GetProducts(ctx)

	var res models.Response

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	res = models.Response{
		Product:    resp,
		Message:    "Retreived Products Successfully",
		StatusCode: 200,
	}

	return res, nil
}
