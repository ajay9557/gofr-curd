package product

import (
	"fmt"
	"gofr-curd/model"
	"gofr-curd/service"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Serv service.Service
}

func New(s service.Service) Handler {
	return Handler{Serv: s}
}

func (h Handler) GetById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	product, err := h.Serv.GetByProductId(id, ctx)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}
	fmt.Println(product)
	return product, nil
}

func (h Handler) DeleteById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	if err := h.Serv.DeleteByProductId(ctx, id); err != nil {
		return nil, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
	}
	return "Deleted Product Successfully", nil
}

func (h Handler) Update(ctx *gofr.Context) (interface{}, error) {
	var product model.ProductDetails
	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	err := h.Serv.UpdateByProductId(ctx, product)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}
	return product, nil
}

func (h Handler) Insert(ctx *gofr.Context) (interface{}, error) {
	var product model.ProductDetails
	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	if product.Id == 0 {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	resp, err := h.Serv.InsertProduct(ctx, product)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}
	return resp, nil
}

func (h Handler) GetAllProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Serv.GetProducts(ctx)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}
	return resp, nil
}
