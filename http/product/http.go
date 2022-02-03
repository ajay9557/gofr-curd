package product

import (
	"fmt"
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
	fmt.Println(product)
	return product, nil
}
