package product

import (
	"gofr-curd/models"
	"gofr-curd/service"
	"gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type services struct {
	store store.Store
}

func New(s store.Store) service.Services {
	return &services{
		store: s,
	}
}

func (s *services) Get(ctx *gofr.Context) ([]*models.Product, error) {
	return s.store.Get(ctx)
}

func (s *services) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	if !validateID(id) {
		return nil, errors.InvalidParam{
			Param: []string{"id"},
		}
	}

	return s.store.GetByID(ctx, id)
}

func (s *services) Create(ctx *gofr.Context, p models.Product) (*models.Product, error) {
	if !validateID(p.ID) {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err := s.store.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	res, err := s.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *services) Update(ctx *gofr.Context, p models.Product) (*models.Product, error) {
	if !validateID(p.ID) {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err := s.store.Update(ctx, p)
	if err != nil {
		return nil, err
	}

	res, err := s.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *services) Delete(ctx *gofr.Context, id int) error {
	if !validateID(id) {
		return errors.InvalidParam{Param: []string{"id"}}
	}

	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
