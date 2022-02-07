package product

import (
	"errors"
	"product/models"
	"product/services"
	"product/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductService struct {
	StoreInterface stores.Store
}

func New(si stores.Store) services.Service {
	return &ProductService{si}
}

func (productService *ProductService) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {
	if id < 0 {
		return models.Product{}, errors.New("INVALID ID")
	}
	product, _ := productService.StoreInterface.GetProductById(ctx, id)
	return product, nil
}

func (productService *ProductService) GetAllProduct(ctx *gofr.Context) ([]models.Product, error) {
	return productService.StoreInterface.GetAllProduct(ctx)
}

func (productService *ProductService) AddProduct(ctx *gofr.Context, product models.Product) error {
	result := productService.StoreInterface.AddProduct(ctx, product)
	if result != nil {
		return errors.New("FAILED TO ADD THE PRODUCT")
	}
	return nil
}

func (productService *ProductService) UpdateProduct(ctx *gofr.Context, product models.Product) error {
	err := productService.StoreInterface.UpdateProduct(ctx, product)
	if err != nil {
		return errors.New("FAILED TO UPDATE THE PRODUCT")
	}
	return nil
}

func (productService *ProductService) DeleteProduct(ctx *gofr.Context, id int) error {
	err := productService.StoreInterface.DeleteProduct(ctx, id)
	if err != nil {
		return errors.New("FAILED TO DELETE THE PRODUCT")
	}
	return nil
}
