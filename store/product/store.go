package product

import (
	"database/sql"
	"gofr-curd/models"
	"gofr-curd/store"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct{}

func New() store.Store {
	return &product{}
}

func (p *product) Get(ctx *gofr.Context) ([]*models.Product, error) {
	var res []*models.Product
	rows, err := ctx.DB().QueryContext(ctx, "select id,name,type from products")
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "products", ID: "all"}
	}

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Type)
		if err != nil {
			return nil, errors.EntityNotFound{Entity: "product"}
		}
		res = append(res, &p)
	}
	return res, nil
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

func (p *product) Create(ctx *gofr.Context, pd models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "insert into products(id,name,type) values(?,?,?)", pd.Id, pd.Name, pd.Type)
	if err != nil {
		return errors.EntityAlreadyExists{}
	}
	return nil
}

func (p *product) Update(ctx *gofr.Context, pd models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "update products set name=?, type=? where id=?", pd.Name, pd.Type, pd.Id)
	if err != nil {
		return errors.Error("error updating record")
	}
	return nil
}

func (p *product) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from products where id=?", id)
	if err != nil {
		return errors.Error("error deleting record")
	}
	return nil
}
