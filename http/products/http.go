package products

import (
	"encoding/json"
	"reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"
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

func (h handler) CreateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	reqBody := ctx.Request().Body

	err := json.NewDecoder(reqBody).Decode(&product)
	if err != nil || reflect.DeepEqual(product, models.Product{}) {
		return nil, errors.MissingParam{Param: []string{"name", "category"}}
	}

	pr, err := h.srv.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Product *models.Product `json:"product"`
	}{
		Product: pr,
	}

	return resData, nil
}

func (h handler) UpdateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	reqBody := ctx.Request().Body

	err := json.NewDecoder(reqBody).Decode(&product)
	if err != nil || reflect.DeepEqual(product, models.Product{}) {
		return nil, errors.MissingParam{Param: []string{"name", "category"}}
	}

	id := ctx.PathParam("id")

	pr, err := h.srv.UpdateById(ctx, id, product)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Product *models.Product `json:"product"`
	}{
		Product: pr,
	}

	return resData, nil
}
