package product

import (
	// "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
	"github.com/tejas/gofr-crud/store"
	"errors"
)

type services struct {
	store1 store.ProductStore
}

func New(s store.ProductStore) service.ProductService {
	return services{
		store1: s,
	}
}

func (s services) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {

	if ValidId(id){
		row,err := s.store1.GetProductById(ctx, id)

		if err != nil {
			return models.Product{}, errors.New("cannot fetch data for the given id")
		}

		return row, nil
	}

	return models.Product{}, errors.New("invalid id")
}
