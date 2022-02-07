/*This file contains the required businees logic implementation. Internally it is using the
  CRUD functionality provided in the stores package*/

package products

import (
	"errors"
	"reflect"

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

//Fetches a product with the given Id
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

//Fetches all the products
func (se *Product) Read(ctx *gofr.Context) ([]models.Product, error) {

	product, err := se.p.Read(ctx)

	if err != nil {
		return nil, err
	}

	return product, nil
}

//Creates a Product entity
func (se *Product) Create(ctx *gofr.Context, value *models.Product) (*models.Product, error) {

	if value == nil || reflect.DeepEqual(value, models.Product{}) {

		return nil, errors.New("Invalid Entity")
	}

	if value.Name == "" || value.Type == "" {

		return nil, errors.New("Invalid name or Type")
	}

	product, err := se.p.Create(ctx, value)

	if err != nil {
		return nil, err
	}

	return product, nil

}

//Updates a product with the given Id
func (se *Product) Update(ctx *gofr.Context, value *models.Product, id int) (*models.Product, error) {

	if value == nil || reflect.DeepEqual(value, models.Product{}) {

		return nil, errors.New("Invalid Entity")
	}

	if id < 1 {

		return nil, errors.New("Invalid Id")
	}

	product, err := se.ReadByID(ctx, id)

	if err != nil || product.Name == "" {

		return nil, errors.New("Invalid Id")
	}

	product, err = se.p.Update(ctx, value, id)

	if err != nil {

		return nil, err
	}

	return product, nil
}

//Deletes a product with the given Id
func (se *Product) Delete(ctx *gofr.Context, id int) error {

	if id < 1 {

		return errors.New("Invalid Id")
	}

	product, err := se.p.ReadByID(ctx, id)

	if err != nil || product.Name == "" {
		return errors.New("Invalid Id")
	}

	err = se.p.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
