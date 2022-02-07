package product

import (
	"gofr-curd/models"
	"gofr-curd/service"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	serv service.Service
}

func New(s service.Service) Handler {
	return Handler{serv: s}
}

func (h Handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	product, err := h.serv.GetByProductID(id, ctx)

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}

	return product, nil
}

func (h Handler) GetAllProductDetails(ctx *gofr.Context) (interface{}, error) {
	allProducts, err := h.serv.GetProducts(ctx)

	if err != nil {
		return nil, errors.Error("internal error")
	}

	return allProducts, nil
}

func (h Handler) InsertProduct(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if err := h.serv.InsertProductDetails(product, ctx); err != nil {
		return nil, errors.Error("internal errror")
	}

	return product, nil
}

func (h Handler) UpdateProductByID(ctx *gofr.Context) (interface{}, error) {
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

func (h Handler) DeleteByProductID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	err = h.serv.DeleteProductByID(id, ctx)

	if err != nil {
		return nil, errors.Error("internal error")
	}

	return "Deleted successfully", nil
}
