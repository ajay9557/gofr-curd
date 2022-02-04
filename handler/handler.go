package handler

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/services"
	_ "github.com/himanshu-kumar-zs/gofr-curd/store"
	"strconv"
)

type handler struct {
	serv services.Product
}

func New(s services.Product) handler {
	return handler{
		serv: s,
	}
}
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
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
