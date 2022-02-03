package products

import (
	"fmt"
	"strconv"
	"zopsmart/gofr-curd/model"
	"zopsmart/gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductStore struct {
	store store.Productstorer
}

func New(productstore store.Productstorer) ProductStore {
	return ProductStore{productstore}
}

func (p ProductStore) GetByID(ctx *gofr.Context, i string) (model.Product, error) {
	isValid, err := validateId(i)
	if !isValid {
		return model.Product{}, err
	}
	id, _ := strconv.Atoi(i)

	resp, err := p.store.GetProductById(ctx, id)
	if err != nil {
		return model.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return resp, nil
}
