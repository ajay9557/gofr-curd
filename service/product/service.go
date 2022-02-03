package product

import (
	"errors"
	"gofr-curd/models"
	"gofr-curd/service"
	"gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ServiceHandler struct {
	store store.Store
}

func New(store store.Store) service.Service {
	return ServiceHandler{store: store}
}

func (s ServiceHandler) GetByProductId(id int, ctx *gofr.Context) (models.Product, error) {
	checkId := idValidation(id)
	if checkId {
		prod, err := s.store.GetById(id, ctx)
		if err != nil {
			return models.Product{}, errors.New("product not found")
		}
		return prod, nil
	}
	return models.Product{}, errors.New("invalid id")
}
