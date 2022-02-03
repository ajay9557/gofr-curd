/*This file contains the required businees logic implementation. Internally it is using the
  CRUD functionality provided in the stores package*/

package products

import (
	"errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/services"
	"github.com/arohanzst/testapp/stores"
)

type Product struct {
	p stores.Product
}

func New(p stores.Product) services.Product {
	return &Product{p}
}

//Fetches a product with the given ID
func (se *Product) ReadByID(ctx *gofr.Context, id int) (*models.Product, error) {

	if id < 1 {

		return nil, errors.New("Invalid Id")
	}

	product, err := se.p.ReadByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
