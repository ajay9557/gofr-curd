package product

import (
	"errors"

	err1 "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
	"github.com/tejas/gofr-crud/store"
)

type services struct {
	store1 store.ProductStore
}

func New(s store.ProductStore) service.ProductService {
	return services{
		store1: s,
	}
}

func (s services) GetProductByID(ctx *gofr.Context, id int) (models.Product, error) {
	if ValidID(id) {
		row, err := s.store1.GetProductByID(ctx, id)

		if err != nil {
			return models.Product{}, errors.New("cannot fetch data for the given id")
		}

		return row, nil
	}

	return models.Product{}, errors.New("invalid id")
}

func (s services) GetAllProducts(ctx *gofr.Context) ([]models.Product, error) {
	rows, err := s.store1.GetAllProducts(ctx)
	if err != nil {
		return nil, errors.New("connot fetch all products data")
	}

	return rows, nil
}

func (s services) UpdateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	if ValidID(prod.ID) {
		updateProd, err := s.store1.UpdateProduct(ctx, prod)

		if err != nil {
			return models.Product{}, errors.New("cannot update the product")
		}

		return updateProd, nil
	}

	return models.Product{}, errors.New("invalid id")
}

func (s services) CreateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	createProduct, err := s.store1.CreateProduct(ctx, prod)

	if err != nil {
		return models.Product{}, errors.New("cannot create the product")
	}

	return createProduct, nil
}

func (s services) DeleteProduct(ctx *gofr.Context, id int) error {
	if !ValidID(id) {
		return err1.Error("error while deleting product")
	}

	err := s.store1.DeleteProduct(ctx, id)

	return err
}
