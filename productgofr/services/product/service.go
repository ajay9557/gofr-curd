package product

import (
	//	"errors"
	models "zopsmart/productgofr/models"
	stores "zopsmart/productgofr/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	//	"github.com/openzipkin/zipkin-go/model"
)

type ProductService struct {
	store stores.Store
}

func New(s stores.Store) *ProductService {
	return &ProductService{
		store: s,
	}
}

func (p ProductService) GetProdByID(ctx *gofr.Context, id int) (models.Product, error) {
	ok := validateID(id)
	if !ok {
		return models.Product{}, errors.InvalidParam{Param: []string{"id"}}
	}
	return p.store.GetProdByID(ctx, id)

}

func (p ProductService) GetAllProd(ctx *gofr.Context) ([]models.Product, error) {

	res, err := p.store.GetAllProduct(ctx)

	if err != nil {
		return []models.Product{}, errors.EntityNotFound{Entity: "products"}
	}

	return res, nil
}

func (p ProductService) DeleteProduct(ctx *gofr.Context, id int) error {

	ok := validateID(id)
	if !ok {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	err := p.store.DeleteProduct(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) UpdateProduct(ctx *gofr.Context, pro models.Product) error {

	ok := validateID(pro.Id)
	if !ok {
		return errors.InvalidParam{Param: []string{"id"}}
	}

	err := p.store.UpdateProduct(ctx, pro)

	if err != nil {
		return errors.EntityNotFound{Entity: "products", ID: "id"}
	}
	return nil

}

func (p ProductService) CreateProduct(ctx *gofr.Context, pro models.Product) error {
	ok := validateID(pro.Id)

	if !ok {
		return errors.InvalidParam{Param: []string{"Id"}}
	}

	err := p.store.CreateProduct(ctx, pro)

	if err != nil {
		return errors.EntityAlreadyExists{}
	}

	return nil
}
