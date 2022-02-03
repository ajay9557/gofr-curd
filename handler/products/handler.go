package products

import (
	"zopsmart/gofr-curd/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Svc service.Productservice
}

func New(svc service.Productservice) Handler {
	return Handler{svc}
}

func (h Handler) GetById(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	resp, err := h.Svc.GetByID(ctx, i)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Product",
			ID:     i,
		}
	}
	return resp, nil

}
