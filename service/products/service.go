package products

import (
	"reflect"
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

func (srv *Service) Create(ctx *gofr.Context, pr models.Product) (*models.Product, error) {
	// Validation of product object (Not allowing empty)
	if reflect.DeepEqual(models.Product{}, pr) {
		return nil, errors.Error("Need Product data to create new product")
	}

	err := srv.store.Create(ctx, pr)
	if err != nil {
		return nil, err
	}

	// Fetch created product
	product, _ := srv.store.GetById(ctx, pr.Id)

	return product, nil
}

func (srv *Service) UpdateById(ctx *gofr.Context, id string, pr models.Product) (*models.Product, error) {

	convId, err := strconv.Atoi(id)
	// Id must be a number
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	// Id must be greater than 0
	if convId < 0 {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	// Checking whether product is exist or not
	_, err = srv.store.GetById(ctx, convId)
	if err != nil {
		return nil, err
	}

	err = srv.store.UpdateById(ctx, convId, pr)
	if err != nil {
		return nil, err
	}

	p, _ := srv.store.GetById(ctx, convId)

	return p, nil
}

func (srv *Service) DeleteById(ctx *gofr.Context, id string) error {
	convId, err := strconv.Atoi(id)
	// Id must be a number
	if err != nil {
		return errors.InvalidParam{Param: []string{"id"}}
	}

	// Id must be greater than 0
	if convId < 0 {
		return errors.InvalidParam{Param: []string{"id"}}
	}

	// Checking whether product is exist or not
	_, err = srv.store.GetById(ctx, convId)
	if err != nil {
		return err
	}

	err = srv.store.DeleteById(ctx, convId)
	if err != nil {
		return err
	}

	return nil
}
