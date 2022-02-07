package product

import (
	"gofr-curd/models"
	"gofr-curd/service"
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

func (h *Handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Data:       *resp,
		Message:    "data retrieved",
		StatusCode: http.StatusOK,
	}, err
}

func (h *Handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.Get(ctx)
	return resp, err
}

func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var p models.Product
	if err := ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.service.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var p models.Product
	p.ID = id

	if err = ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.service.Update(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *Handler) Delete(ctx *gofr.Context) (interface{}, error) {
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
