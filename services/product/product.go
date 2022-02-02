package product

import (
	"errors"
	"product/models"
	"product/services"
	"product/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductService struct {
	storeInterface stores.Store
}

func New(si stores.Store) services.Service {
	return &ProductService{si}
}

func (productService *ProductService) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {
	if id < 0 {
		return models.Product{}, errors.New("INVALID ID")
	}
	product, _ := productService.storeInterface.GetProductById(ctx, id)
	return product, nil
}
