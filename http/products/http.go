package products

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
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
