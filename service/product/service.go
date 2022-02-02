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

func (s *services) GetById(ctx *gofr.Context, id int) (*models.Product, error) {
	if !validateId(id) {
		return nil, errors.InvalidParam{
			Param: []string{"id"},
		}
	}
	return s.store.GetById(ctx, id)
}
