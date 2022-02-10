package product

import (
	//	"errors"
	models "zopsmart/productgofr/models"
	stores "zopsmart/productgofr/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
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

func (p ProductService) GetProdByID(ctx *gofr.Context, id int) (models.Product, error) {
	ok := validateID(id)
	if !ok {
		return models.Product{},errors.InvalidParam{Param: []string{"id"}}
	}
	res, err := p.store.GetProdByID(ctx, id)

	if err != nil {
		return res, err
	}

	return res, nil

}

func (p ProductService) GetAllProd(ctx *gofr.Context) ([]models.Product, error) {
	return p.store.GetAllProduct(ctx)
}

func (p ProductService) DeleteProduct(ctx *gofr.Context, id int) error {


	err := p.store.DeleteProduct(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) UpdateProduct(ctx *gofr.Context, pro models.Product) error {

	err := p.store.UpdateProduct(ctx, pro)

	if err != nil {
		return errors.InvalidParam{Param: []string{"id"}}
	}
	return nil

}

func (p ProductService) CreateProduct(ctx *gofr.Context, pro models.Product) error {
	// er := check(models.Product)
	// if er!=nil {
	// 	return er
	// }
	err := p.store.CreateProduct(ctx, pro)

	if err != nil {
		return err
	}
	return nil
}
