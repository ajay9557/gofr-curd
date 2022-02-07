package product

import (
	"database/sql"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/store"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

type dbStore struct{}

func New() store.ProductStore {
	return dbStore{}
}

func (p dbStore) GetProductByID(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, "select id, name, type from product where id = ?", id).Scan(&product.ID, &product.Name, &product.Type)

	if err == sql.ErrNoRows {
		return models.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}

	return product, nil
}

func (p dbStore) GetAllProducts(ctx *gofr.Context) ([]models.Product, error) {
	rows, err := ctx.DB().QueryContext(ctx, "select * from product")

	if err != nil {
		return nil, errors.Error("internal query error")
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	products := make([]models.Product, 0)

	for rows.Next() {
		var prod models.Product

		err := rows.Scan(&prod.ID, &prod.Name, &prod.Type)

		if err != nil {
			return nil, err
		}

		products = append(products, prod)
	}

	return products, nil
}

func (p dbStore) UpdateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "update product set name = ?, type = ? where id = ?", prod.Name, prod.Type, prod.ID)

	if err != nil {
		return models.Product{}, errors.DB{Err: err}
	}

	return prod, nil
}

func (p dbStore) CreateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "insert into product values (?,?,?)", prod.ID, prod.Name, prod.Type)

	if err != nil {
		return prod, errors.Error("error in inserting new product")
	}

	return prod, nil
}

func (p dbStore) DeleteProduct(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from product where id = ?", id)

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
