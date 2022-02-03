package products

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/service"
)

type handler struct {
	srv service.Product
}

func New(s service.Product) handler {
	return handler{
		srv: s,
	}
}

func (h handler) GetByIdHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")

	p, err := h.srv.GetById(ctx, param)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Product *models.Product `json:"product"`
	}{
		Product: p,
	}

	return resData, nil
}

func (h handler) GetHandler(ctx *gofr.Context) (interface{}, error) {
	products, err := h.srv.Get(ctx)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Products []*models.Product `json:"products"`
	}{
		Products: products,
	}

	return resData, nil
}
