package product

import (
	"database/sql"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

type Store struct{}

func New() Store {
	return Store{}
}
func (ps Store) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	var resp models.Product

	err := ctx.DB().QueryRowContext(ctx, "select * from product where id = ?", id).Scan(&resp.ID, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return &resp, nil
}
func (ps Store) Update(ctx *gofr.Context, product *models.Product) error {
	feilds, values := buildQuery(product)
	_, err := ctx.DB().ExecContext(ctx, "update product set "+feilds+" where id = ?", values...)

	return err
}

func (ps Store) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM product where id = ?", id)
	return err
}

func (ps Store) Create(ctx *gofr.Context, product *models.Product) (*models.Product, error) {
	res, err := ctx.DB().ExecContext(ctx, "insert into product (name, type) values(?,?)", product.Name, product.Type)

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	id, _ := res.LastInsertId()
	product.ID = int(id)

	return product, nil
}

func (ps Store) GetAll(ctx *gofr.Context) ([]*models.Product, error) {
	var resp []*models.Product

	rows, err := ctx.DB().QueryContext(ctx, "select id, name, type from product")
	if err == sql.ErrNoRows {
		return []*models.Product{}, errors.EntityNotFound{
			Entity: "product",
		}
	}

	for rows.Next() {
		var prod models.Product
		_ = rows.Scan(&prod.ID, &prod.Name, &prod.Type)
		resp = append(resp, &prod)
	}

	return resp, nil
}
