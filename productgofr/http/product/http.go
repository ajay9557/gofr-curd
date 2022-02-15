package product

import (
	models "zopsmart/productgofr/models"
	service "zopsmart/productgofr/services"
	"net/http"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	service service.Services
}

func New(s service.Services) Handler {
	return Handler{
		service: s,
	}
}

func (h *Handler) GetProdByIdHandler(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.service.GetProdByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.Response{
		Data:       resp,
		Message:    "data retrieved",
		StatusCode: http.StatusOK,
	}, err
}

func (h *Handler) GetAllProductHandler(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.GetAllProd(ctx)
	if err!= nil {
		return nil,errors.EntityNotFound{Entity: "products", ID: "all"}
	}
	return resp, nil
}

func (h *Handler) CreateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var p models.Product
	if err := ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	err := h.service.CreateProduct(ctx, p)
	if err != nil {
		return nil, errors.EntityAlreadyExists{}
	}

	return "Successfully created", nil
}

func (h *Handler) UpdateProductHandler(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var p models.Product
	p.Id = id

	if err = ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	er := h.service.UpdateProduct(ctx, p)
	if er != nil {
		return nil, errors.Error("error updating record")
	}

	return "Successfully updated",nil
}

func (h *Handler) DeleteProductHandler(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = h.service.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Message:    "deleted successfully",
		StatusCode: http.StatusOK,
	}, nil
}
