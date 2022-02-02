package product

import (
	"database/sql"
	"gofrPractice/models"
	"gofrPractice/store"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct{}

func New() store.Store {
	return &product{}
}

func (p *product) GetById(ctx *gofr.Context, id int) (*models.Product, error) {
	var resp models.Product
	resp.Id = id
	row := ctx.DB().QueryRowContext(ctx, "select name, type from products where id=?", id)
	err := row.Scan(&resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: strconv.Itoa(id)}
	}

	return &resp, nil
}
