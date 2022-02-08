package products

import (
	"encoding/json"
	"net/http"
	"reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/service"
)

type Handler struct {
	srv service.Product
}

func New(s service.Product) Handler {
	return Handler{
		srv: s,
	}
}

func (h Handler) GetByIDHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")

	p, err := h.srv.GetByID(ctx, param)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Product    *models.Product `json:"product"`
		StatusCode int             `json:"statusCode"`
		Message    string          `json:"message"`
	}{
		Product:    p,
		StatusCode: http.StatusOK,
		Message:    "Successful operation",
	}

	return resData, nil
}

func (h Handler) GetHandler(ctx *gofr.Context) (interface{}, error) {
	products, err := h.srv.Get(ctx)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Products   []*models.Product `json:"products"`
		StatusCode int               `json:"statusCode"`
		Message    string            `json:"message"`
	}{
		Products:   products,
		StatusCode: http.StatusOK,
		Message:    "Successful operation",
	}

	return resData, nil
}

func (h Handler) CreateProductHandler(ctx *gofr.Context) (interface{}, error) {
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
		Product    *models.Product `json:"product"`
		StatusCode int             `json:"statusCode"`
		Message    string          `json:"message"`
	}{
		Product:    pr,
		StatusCode: http.StatusCreated,
		Message:    "Successful operation",
	}

	return resData, nil
}

func (h Handler) UpdateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	reqBody := ctx.Request().Body

	err := json.NewDecoder(reqBody).Decode(&product)
	if err != nil || reflect.DeepEqual(product, models.Product{}) {
		return nil, errors.MissingParam{Param: []string{"name", "category"}}
	}

	id := ctx.PathParam("id")

	pr, err := h.srv.UpdateByID(ctx, id, product)
	if err != nil {
		return nil, err
	}

	resData := &struct {
		Product    *models.Product `json:"product"`
		StatusCode int             `json:"statusCode"`
		Message    string          `json:"message"`
	}{
		Product:    pr,
		StatusCode: http.StatusOK,
		Message:    "Successful operation",
	}

	return resData, nil
}

func (h Handler) DeleteProductHandler(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	err := h.srv.DeleteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
