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
			return product, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
		}
		return res, nil
	}
	return product, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
}

func (pd ProductDetails) DeleteByProductId(ctx *gofr.Context, i int) error {
	Idcheck := checkId(i)
	if Idcheck {
		err := pd.store.DeleteId(ctx, i)
		if err != nil {
			return errors.Error("Internal DB error")
		}
		return nil
	}
	return errors.InvalidParam{Param: []string{"id"}}
}

func (pd ProductDetails) UpdateByProductId(ctx *gofr.Context, product models.Product) error {
	Idcheck := checkId(product.Id)
	if Idcheck {
		err := pd.store.UpdateId(ctx, product)
		if err != nil {
			return errors.Error("Internal DB error")
		}
		return nil
	}
	return errors.InvalidParam{Param: []string{"id"}}
}

func (pd ProductDetails) InsertProduct(ctx *gofr.Context, product models.Product) (models.Product, error) {
	var newProduct models.Product

	newProduct.Id = product.Id
	newProduct.Name = product.Name
	newProduct.Type = product.Type
	Idcheck := checkId(product.Id)
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
