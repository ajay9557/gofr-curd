package product

import (
	"database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"fmt"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

type productStore struct{}

func New() productStore {
	return productStore{}
}
func (ps productStore) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	var resp models.Product
	err := ctx.DB().QueryRowContext(ctx, "select * from product where id = ?", id).Scan(&resp.Id, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return &resp, nil
}
