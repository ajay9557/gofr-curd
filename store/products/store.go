package products

import (
	"database/sql"
	"fmt"
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type DbStore struct {
}

func New() DbStore {
	return DbStore{}
}

func (s *DbStore) GetProductById(ctx *gofr.Context, id int) (model.Product, error) {
	var resp model.Product
	err := ctx.DB().QueryRowContext(ctx, "Select * from Products where id=?", id).Scan(&resp.Id, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return model.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return resp, nil

}
