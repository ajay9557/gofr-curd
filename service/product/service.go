package product

import (
	"fmt"
	"gofr-curd/models"
	"gofr-curd/service"
	"gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"
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
			return models.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
		}
		return prod, nil
	}
	return models.Product{}, errors.InvalidParam{Param: []string{"id"}}
}

func (s ServiceHandler) GetProducts(ctx *gofr.Context) ([]models.Product, error) {
	var allProducts []models.Product
	allProducts, err := s.store.GetAllProducts(ctx)
	if err != nil {
		return nil, errors.Error("Error in database")
	}
	return allProducts, nil
}

func (s ServiceHandler) InsertProductDetails(product models.Product, ctx *gofr.Context) error {
	checkId := idValidation(product.Id)
	if !checkId {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	err := s.store.InsertProduct(product, ctx)
	if err != nil {
		return errors.Error("Error in database")
	}
	return nil
}

func (s ServiceHandler) UpdateProductDetails(product models.Product, ctx *gofr.Context) error {
	checkId := idValidation(product.Id)
	if !checkId {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	err := s.store.UpdateProduct(product, ctx)
	if err != nil {
		return errors.Error("Error in database")
	}
	return nil
}

func (s ServiceHandler) DeleteProductById(id int, ctx *gofr.Context) error {
	checkId := idValidation(id)
	if !checkId {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	err := s.store.DeleteById(id, ctx)
	if err != nil {
		return errors.Error("Error in database")
	}
	return nil
}
