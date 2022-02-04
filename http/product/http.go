package product

import (
	"gofr-curd/models"
	"gofr-curd/service"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type handler struct {
	serv service.Service
}

func New(s service.Service) handler {
	return handler{serv: s}
}

func (h handler) GetById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	product, err := h.serv.GetByProductId(id, ctx)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}
	return product, nil
}

func (h handler) GetAllProductDetails(ctx *gofr.Context) (interface{}, error) {
	allProducts, err := h.serv.GetProducts(ctx)
	if err != nil {
		return nil, errors.Error("internal error")
	}
	return allProducts, nil
}

func (h handler) InsertProduct(ctx *gofr.Context) (interface{}, error) {
	var product models.Product
	err := ctx.Bind(&product)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	err = h.serv.InsertProductDetails(product, ctx)
	if err != nil {
		return nil, errors.Error("internal errror")
	}
	return product, nil
}

func (h handler) UpdateProductById(ctx *gofr.Context) (interface{}, error) {
	var product models.Product
	err := ctx.Bind(&product)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	err = h.serv.UpdateProductDetails(product, ctx)
	if err != nil {
		return nil, errors.Error("internal errror")
	}
	return product, nil
}

func (h handler) DeleteByProductId(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	err = h.serv.DeleteProductById(id, ctx)
	if err != nil {
		return nil, errors.Error("internal error")
	}
	return "Deleted successfully", nil
}
