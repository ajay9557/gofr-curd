package product

import (
	"database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"fmt"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

type ProductStore struct{}

func New() ProductStore {
	return ProductStore{}
}
func (ps ProductStore) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	var resp models.Product
	err := ctx.DB().QueryRowContext(ctx, "select * from product where id = ?", id).Scan(&resp.ID, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return &resp, nil
}
func (ps ProductStore) Update(ctx *gofr.Context, product *models.Product) error {
	feilds, values := buildQuery(product)
	_, err := ctx.DB().ExecContext(ctx, "update product set "+feilds+" where id = ?", values...)
	return err
}
func (ps ProductStore) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM product where id = ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
func (ps ProductStore) Create(ctx *gofr.Context, product *models.Product) (*models.Product, error) {

	res, err := ctx.DB().ExecContext(ctx, "insert into product (name, type) values(?,?)", product.Name, product.Type)

	if err != nil {
		return nil, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	product.ID = int(id)
	return product, nil
}
func (ps ProductStore) GetAll(ctx *gofr.Context) ([]*models.Product, error) {
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
