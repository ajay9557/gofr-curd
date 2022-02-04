package products

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/services"
)

type Handler struct {
	S services.Product
}

func New(s services.Product) Handler {
	return Handler{
		S: s,
	}
}

/*
URL: /product/{id}
Method: GET
Description: Retrieves product with the given ID
*/
func (h Handler) ReadByIdHandler(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	// params := mux.Vars(req)

	// productId := params["id"]

	resp, err := h.S.ReadByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (h Handler) ReadHandler(ctx *gofr.Context) (interface{}, error) {

	resp, err := h.S.Read(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h Handler) CreateHandler(ctx *gofr.Context) (interface{}, error) {

	var p models.Product

	if err := ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	// if p.ID != 0 {
	// 	return nil, errors.InvalidParam{Param: []string{"id"}}
	// }
	ctx.Log("INFO", p)
	resp, err := h.S.Create(ctx, &p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h Handler) UpdateHandler(ctx *gofr.Context) (interface{}, error) {

	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var p models.Product
	if err = ctx.Bind(&p); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	p.Id = id

	resp, err := h.S.Update(ctx, &p, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h Handler) DeleteHandler(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.S.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}
