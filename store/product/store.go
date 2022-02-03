package product

import (
	"database/sql"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type productStorer struct {
}

func New() store.Store {
	return productStorer{}
}

func (p productStorer) GetById(id int, ctx *gofr.Context) (models.Product, error) {

	ReadQ := "Select Id,Name,Type from product where Id=?"
	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, ReadQ, id).Scan(&product.Id, &product.Name, &product.Type)
	if err == sql.ErrNoRows {
		return models.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return product, nil
}
