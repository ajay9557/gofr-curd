package product

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) Service {
	return Service{
		store: s,
	}
}

func (ps Service) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	if !validateID(id) {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	return ps.store.GetByID(ctx, id)
}

func (ps Service) Update(ctx *gofr.Context, product *models.Product) (*models.Product, error) {
	// check if id exists or not
	id := product.ID

	_, err := ps.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// try to update
	err = ps.store.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	updatedProduct, _ := ps.GetByID(ctx, id)

	return updatedProduct, nil
}

func (ps Service) Delete(ctx *gofr.Context, id int) error {
	if !validateID(id) {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	// check if id exists
	_, err := ps.GetByID(ctx, id)
	if err != nil {
		return errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return ps.store.Delete(ctx, id)
}

func (ps Service) Create(ctx *gofr.Context, product *models.Product) (*models.Product, error) {
	if product.Name == "" {
		return nil, errors.MissingParam{Param: []string{"name"}}
	}

	if product.Type == "" {
		return nil, errors.MissingParam{Param: []string{"type"}}
	}

	return ps.store.Create(ctx, product)
}

func (ps Service) GetAll(ctx *gofr.Context) ([]*models.Product, error) {
	return ps.store.GetAll(ctx)
}
