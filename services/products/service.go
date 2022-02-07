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

func (pd ProductDetails) GetByUserID(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product

	Idcheck := checkID(id)

	if Idcheck {
		res, err := pd.store.GetID(ctx, id)

		if err != nil {
			return product, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
		}

		return res, nil
	}

	return product, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
}

func (pd ProductDetails) DeleteByProductID(ctx *gofr.Context, i int) error {
	Idcheck := checkID(i)

	if Idcheck {
		err := pd.store.DeleteID(ctx, i)

		if err != nil {
			return errors.Error("Internal DB error")
		}

		return nil
	}

	return errors.InvalidParam{Param: []string{"id"}}
}

func (pd ProductDetails) UpdateByProductID(ctx *gofr.Context, product models.Product) error {
	Idcheck := checkID(product.ID)

	if Idcheck {
		err := pd.store.UpdateID(ctx, product)

		if err != nil {
			return errors.Error("Internal DB error")
		}

		return nil
	}

	return errors.InvalidParam{Param: []string{"id"}}
}

func (pd ProductDetails) InsertProduct(ctx *gofr.Context, product models.Product) (models.Product, error) {
	var newProduct models.Product

	newProduct.ID = product.ID
	newProduct.Name = product.Name
	newProduct.Type = product.Type
	Idcheck := checkID(product.ID)

	if Idcheck {
		_, err := pd.store.CreateProducts(ctx, product)

		if err != nil {
			return newProduct, errors.Error("Internal DB error")
		}

		return newProduct, nil
	}

	return newProduct, errors.InvalidParam{Param: []string{"id"}}
}

func (pd ProductDetails) GetProducts(ctx *gofr.Context) ([]models.Product, error) {
	var products []models.Product

	res, err := pd.store.GetAll(ctx)

	if err != nil {
		return products, errors.Error("Internal DB error")
	}

	return res, nil
}
