package products

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/service"
	"github.com/ridhdhish-desai-zs/product-gofr/store"
)

type Service struct {
	store store.Product
}

func New(s store.Product) service.Product {
	return &Service{
		store: s,
	}
}

func (srv *Service) GetById(ctx *gofr.Context, id string) (*models.Product, error) {
	convId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "products", ID: id}
	}

	if convId < 0 {
		return nil, errors.EntityNotFound{Entity: "products", ID: id}
	}

	product, err := srv.store.GetById(ctx, convId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (srv *Service) Get(ctx *gofr.Context) ([]*models.Product, error) {
	products, err := srv.store.Get(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
