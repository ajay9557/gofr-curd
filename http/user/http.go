package user

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Services"
)

type Handler struct {
	ser Services.Serviceint
}

func New(s Services.Serviceint) Handler {
	return Handler{ser: s}
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

	resp, err := h.ser.GetId(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			ID: i,
		}
	}

	return resp, nil
}
