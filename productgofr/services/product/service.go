package product

import (
	models "zopsmart/productgofr/models"
	stores "zopsmart/productgofr/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductService struct {
	store stores.Store
}

func New(s stores.Store) *ProductService {
	return &ProductService{
		store: s,
	}
}

func (p ProductService) GetProdByID(ctx *gofr.Context, id int) (*models.Product, error) {
	res, err := p.store.GetProdByID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil

}
