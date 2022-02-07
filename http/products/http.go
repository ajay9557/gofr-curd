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

	return resp, nil
}

func (h Handler) DeleteByID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

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

	return "Deleted Product Successfully", nil
}

func (h Handler) Update(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	err := h.Service.UpdateByProductID(ctx, product)

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	return product, nil
}

func (h Handler) Insert(ctx *gofr.Context) (interface{}, error) {
	var product models.Product
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

	return resp, nil
}

func (h Handler) GetAllProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Service.GetProducts(ctx)

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	return resp, nil
}
