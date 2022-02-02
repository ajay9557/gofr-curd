package products

import (
	"database/sql"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/store"
)

type product struct{}

func New() store.Product {
	return product{}
}

func (p product) GetById(ctx *gofr.Context, id int) (*models.Product, error) {
	var product models.Product

	err := ctx.DB().QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.Id, &product.Name, &product.Category)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return &product, nil
}
