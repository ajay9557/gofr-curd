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

func (p product) GetByID(ctx *gofr.Context, id int) (*models.Product, error) {
	row := ctx.DB().QueryRowContext(ctx, "SELECT * FROM products WHERE id = ?", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Category)

	if err != nil || err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	return &product, nil
}

func (p product) Get(ctx *gofr.Context) ([]*models.Product, error) {
	var products []*models.Product

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pr models.Product
		err := rows.Scan(&pr.ID, &pr.Name, &pr.Category)

		if err != nil {
			return nil, errors.EntityNotFound{Entity: "product"}
		}

		products = append(products, &pr)
	}

	if products == nil {
		return nil, errors.EntityNotFound{Entity: "product"}
	}

	return products, nil
}

func (p product) Create(ctx *gofr.Context, pr models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO products(id, name, category) values(?, ?, ?)", pr.ID, pr.Name, pr.Category)
	if err != nil {
		return errors.Error("Connection lost")
	}

	return nil
}

func (p product) UpdateByID(ctx *gofr.Context, id int, pr models.Product) error {
	query := "UPDATE products SET"

	fields, args := formUpdateQuery(pr)
	args = append(args, id)

	subQuery := fields[:len(fields)-1]
	query += subQuery + " WHERE id = ?"

	_, err := ctx.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Error("Connection lost")
	}

	return nil
}

func (p product) DeleteByID(ctx *gofr.Context, id int) error {
	result, err := ctx.DB().ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return nil
}
