package product

import (
	"gofr-curd/models"
	"gofr-curd/service"
	"net/http"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type handler struct {
	service service.Services
}

func New(s service.Services) handler {
	return handler{
		service: s,
	}
}

func (h *handler) GetById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.service.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Response{
		Data:       *resp,
		Message:    "data retrieved",
		StatusCode: http.StatusOK,
	}, err
}

func (h *handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.Get(ctx)
	return resp, err
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var p models.Product
	if err := ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if p.Id != 0 {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.service.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (h *handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var p models.Product
	if err := ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	p.Id = id
	resp, err := h.service.Update(ctx, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = h.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Response{
		Message:    "deleted successfully",
		StatusCode: http.StatusOK,
	}, nil
}
