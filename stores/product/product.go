package product

import (
	"product/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store struct {
}

func New() Store {
	return Store{}
}

func (store Store) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {
	product := models.Product{}
	err := ctx.DB().QueryRow("select * from product where id=?", id).Scan(&product.Id, &product.Name, &product.Type)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
