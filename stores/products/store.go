package products

import (
	"database/sql"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
)

type ProductStorer struct {
}

func New() stores.Store {
	return &ProductStorer{}
}

func (p *ProductStorer) GetId(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product
	err := ctx.DB().QueryRowContext(ctx, "Select Id,Name,Type from Product where Id =?", id).Scan(&product.Id, &product.Name, &product.Type)
	if err == sql.ErrNoRows {
		return product, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
	}
	return product, nil
}
