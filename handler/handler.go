package handler

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/services"
)

type Handler struct {
	serv services.Product
}

func New(s services.Product) Handler {
	return Handler{
		serv: s,
	}
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

	resp, err := h.serv.GetByID(ctx, id)

	return resp, err
}

func (h Handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var newProduct models.Product
	if err = ctx.Bind(&newProduct); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	newProduct.ID = id
	resp, err := h.serv.Update(ctx, &newProduct)

	return resp, err
}

func (h Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = h.serv.Delete(ctx, id)

	return "successfully deleted", err
}

func (h Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var newProduct models.Product
	if err := ctx.Bind(&newProduct); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.serv.Create(ctx, &newProduct)

	return resp, err
}

func (h Handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	return h.serv.GetAll(ctx)
}
