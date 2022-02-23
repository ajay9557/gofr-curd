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
	isValid, err := validateID(i)
	if !isValid {
		return model.Product{}, err
	}

	id, _ := strconv.Atoi(i)

	resp, err := p.store.GetProductByID(ctx, id)
	if err != nil {
		return model.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return resp, nil
}

func (p ProductStore) DeleteByID(ctx *gofr.Context, i string) error {
	isValid, err := validateID(i)
	if !isValid {
		return err
	}

	id, _ := strconv.Atoi(i)

	err = p.store.DeleteByID(ctx, id)
	if err != nil {
		return errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return nil
}

func (p ProductStore) GetProducts(ctx *gofr.Context) ([]model.Product, error) {
	resp, err := p.store.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p ProductStore) AddProduct(ctx *gofr.Context, prod model.Product) (model.Product, error) {
	isValid, err := CheckBody(prod)
	if !isValid {
		return model.Product{}, err
	}

	id, err := p.store.AddProduct(ctx, prod)
	if err != nil {
		return model.Product{}, err
	}

	prod.ID = id

	return prod, nil
}

func (p ProductStore) UpdateByID(ctx *gofr.Context, prod model.Product, id string) (model.Product, error) {
	isValid, err := validateID(id)
	if !isValid {
		return model.Product{}, err
	}

	isValid, err = validate(prod.ID)
	if !isValid {
		return model.Product{}, err
	}

	i, _ := strconv.Atoi(id)
	prod.ID = i

	resp, err := p.store.UpdateByID(ctx, prod)
	if err != nil {
		return model.Product{}, errors.EntityNotFound{Entity: "product", ID: id}
	}

	return resp, nil
}
