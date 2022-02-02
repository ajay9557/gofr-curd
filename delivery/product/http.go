package product

import (
	"gofr-curd/service"
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
	return resp, err
}
