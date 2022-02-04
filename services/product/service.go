package product

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/store"
)

type productService struct {
	store store.Store
}

func New(s store.Store) productService {
	return productService{
		store: s,
	}
}

func (ps productService) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	if !validateID(id) {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	return ps.store.GetByID(ctx, id)
}
