package user

import (
	"fmt"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Services"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
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

func (h Handler) Create(ctx *gofr.Context) (interface{}, error) {
	prod := model.Product{}
	err := ctx.Bind(&prod)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.ser.Create(ctx, prod)
	fmt.Println(resp)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	return resp, nil
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

	err = h.ser.Delete(id, ctx)
	if err != nil {
		return nil, errors.EntityNotFound{
			ID: i,
		}
	}

	return nil, nil
}

func (h Handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.ser.GetUser(ctx)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	return resp, nil
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

	prod := model.Product{}
	err = ctx.Bind(&prod)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	prod.Id = id

	resp, err := h.ser.Update(prod, ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
