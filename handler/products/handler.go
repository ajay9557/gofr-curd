package products

import (
	"fmt"
	"zopsmart/gofr-curd/model"
	"zopsmart/gofr-curd/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Svc service.Productservice
}

func New(svc service.Productservice) Handler {
	return Handler{svc}
}

func (h Handler) GetById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	resp, err := h.Svc.GetByID(ctx, i)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Product",
			ID:     i,
		}
	}
	return resp, nil

}

func (h Handler) UpdateById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	var prod model.Product
	if err := ctx.Bind(&prod); err != nil {
		ctx.Logger.Errorf("error in binding : %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := h.Svc.UpdateById(ctx, prod, i)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (h Handler) GetProducts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Svc.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h Handler) AddProduct(ctx *gofr.Context) (interface{}, error) {
	var prod model.Product
	if err := ctx.Bind(&prod); err != nil {
		ctx.Logger.Errorf("error in binding : %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := h.Svc.AddProduct(ctx, prod)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (h Handler) DeleteById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	err := h.Svc.DeleteById(ctx, i)
	fmt.Printf("%T", i)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Product",
			ID:     i,
		}
	}
	return "Deleted Successfully", nil

}
