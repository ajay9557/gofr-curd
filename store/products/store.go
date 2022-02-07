package products

import (
	"database/sql"
	"fmt"
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type DBStore struct {
}

func New() DBStore {
	return DBStore{}
}

func (s *DBStore) GetProductByID(ctx *gofr.Context, id int) (model.Product, error) {
	var resp model.Product

	err := ctx.DB().QueryRowContext(ctx, "Select * from Products where id=?", id).Scan(&resp.ID, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		return model.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return resp, nil
}

func (s *DBStore) DeleteByID(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "Delete from Products where id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s *DBStore) GetProducts(ctx *gofr.Context) ([]model.Product, error) {
	rows, err := ctx.DB().QueryContext(ctx, "Select * from Products")
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var pd model.Product

		err := rows.Scan(&pd.ID, &pd.Name, &pd.Type)
		if err != nil {
			return nil, err
		}

		products = append(products, pd)
	}

	return products, nil
}

func (s *DBStore) AddProduct(ctx *gofr.Context, prod model.Product) (int, error) {
	res, err := ctx.DB().ExecContext(ctx, "INSERT INTO Products(Id,Name,Type) VALUES(?,?,?)", prod.ID, prod.Name, prod.Type)
	if err != nil {
		return -1, errors.DB{Err: err}
	}

	lastID, _ := res.LastInsertId()

	return int(lastID), nil
}

func (s *DBStore) UpdateByID(ctx *gofr.Context, prod model.Product) (model.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "Update Products set Name=?,Type=? where Id=?", prod.Name, prod.Type, prod.ID)
	if err != nil {
		return model.Product{}, errors.DB{Err: err}
	}

	return prod, nil
}
