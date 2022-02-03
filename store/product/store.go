package product

import (
	"database/sql"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

type dbStore struct {
}

func New() store.ProductStore {
	return dbStore{}
}

func (p dbStore) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {

	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, "select id, name, type from product where id = ?", id).Scan(&product.Id, &product.Name, &product.Type)

	if err == sql.ErrNoRows {
		return models.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return product, nil
}
