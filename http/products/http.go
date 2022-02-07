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

	resp, err := h.S.ReadByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

/*
URL: /product
Method: GET
Description: Retrieves all the existing products in the product database
*/
func (h Handler) ReadHandler(ctx *gofr.Context) (interface{}, error) {

	resp, err := h.S.Read(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
URL: /product
Method: POST
Description: Creates a product entity with the given name and type
*/
func (h Handler) CreateHandler(ctx *gofr.Context) (interface{}, error) {

	var p models.Product

	err := ctx.Bind(&p)
	if err != nil {

		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	ctx.Log("INFO", p)
	resp, err := h.S.Create(ctx, &p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
URL: /product/{id}
Method: PUT
Description: Updates a product with the given ID
*/
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
	err = ctx.Bind(&p)
	if err != nil {
		return nil, errors.MissingParam{Param: []string{"Name", "Type"}}
	}

	p.Id = id

	resp, err := h.S.Update(ctx, &p, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
URL: /product/{id}
Method: DELETE
Description: Deletes a product with the given ID
*/
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
