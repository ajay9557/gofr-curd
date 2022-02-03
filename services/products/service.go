package products

import (
	"fmt"
	"gofr-curd/models"
	"gofr-curd/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductDetails struct {
	store stores.Store
}

func New(ps stores.Store) ProductDetails {
	return ProductDetails{ps}
}

func (pd ProductDetails) GetByUserId(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product
	Idcheck := checkId(id)
	if Idcheck {
		res, err := pd.store.GetId(ctx, id)
		if err != nil {
			return product, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
		}
		return res, nil
	}
	return product, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
}
