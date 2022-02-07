package product

import (
	"fmt"
	"gofr-curd/model"
	"gofr-curd/service"
	store "gofr-curd/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ServiceHandler struct {
	store store.Store
}

func New(store store.Store) service.Service {
	return ServiceHandler{store: store}
}

func (s ServiceHandler) GetByProductId(id int, ctx *gofr.Context) (model.ProductDetails, error) {
	checkId := idValidation(id)
	if checkId {
		prod, err := s.store.GetProductById(id, ctx)
		if err != nil {
			return model.ProductDetails{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
		}
		return prod, nil
	}
	return model.ProductDetails{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
}

func (s ServiceHandler) DeleteByProductId(ctx *gofr.Context, id int) error {
	Idcheck := idValidation(id)
	if Idcheck {
		err := s.store.DeleteProductId(ctx, id)
		if err != nil {
			return errors.Error("Internal Database Error")
		}
		return nil
	}
	return errors.InvalidParam{Param: []string{"id"}}
}
func (s ServiceHandler) UpdateByProductId(ctx *gofr.Context, product model.ProductDetails) error {
	Idcheck := idValidation(product.Id)
	if Idcheck {
		err := s.store.UpdateProductById(ctx, product)
		if err != nil {
			return errors.Error("Internal DB error")
		}
		return nil
	}
	return errors.InvalidParam{Param: []string{"id"}}
}
func (s ServiceHandler) InsertProduct(ctx *gofr.Context, product model.ProductDetails) (model.ProductDetails, error) {
	var newProduct model.ProductDetails

	newProduct.Id = product.Id
	newProduct.Name = product.Name
	newProduct.Types = product.Types
	Idcheck := idValidation(product.Id)
	if Idcheck {
		_, err := s.store.CreateProducts(ctx, product)
		if err != nil {
			return newProduct, errors.Error("Internal DB error")
		}
		return newProduct, nil
	}
	return newProduct, errors.InvalidParam{Param: []string{"id"}}
}

func (s ServiceHandler) GetProducts(ctx *gofr.Context) ([]model.ProductDetails, error) {
	var products []model.ProductDetails
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return products, errors.Error("Internal DB error")
	}
	return res, nil
}
