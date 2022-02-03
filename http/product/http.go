package product

import (
	"fmt"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/service"
)

type handler struct {
	service1 service.ProductService
}

func New(s service.ProductService) handler {
	return handler{
		service1: s,
	}
}

func (h handler) GetProductById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return id, errors.InvalidParam{Param: []string{"id"}}
	}

	res, err := h.service1.GetProductById(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     i,
		}
	}
	fmt.Println(res)
	return res, nil
}
