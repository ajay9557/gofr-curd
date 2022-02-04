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

func (p product) GetById(ctx *gofr.Context, id int) (*models.Product, error) {
	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM products WHERE id = ?", id).Scan(&product.Id, &product.Name, &product.Category)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return &product, nil
}

func (p product) Get(ctx *gofr.Context) ([]*models.Product, error) {

	var products []*models.Product

	rows, _ := ctx.DB().QueryContext(ctx, "SELECT * FROM products")

	for rows.Next() {
		var pr models.Product
		err := rows.Scan(&pr.Id, &pr.Name, &pr.Category)
		if err != nil {
			return nil, errors.EntityNotFound{Entity: "product"}
		}

		products = append(products, &pr)
	}

	return products, nil
}

func (p product) Create(ctx *gofr.Context, pr models.Product) (int, error) {
	result, _ := ctx.DB().ExecContext(ctx, "INSERT INTO products(name, category) values(?, ?)", pr.Name, pr.Category)

	id, _ := result.LastInsertId()

	return int(id), nil
}

func (p product) UpdateById(ctx *gofr.Context, id int, pr models.Product) error {

	query := "UPDATE products SET"

	fields, args := formUpdateQuery(pr)
	args = append(args, id)

	subQuery := fields[:len(fields)-1]
	query += subQuery + "WHERE id = ?"

	_, err := ctx.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Error("Connection lost")
	}

	return nil
}

func (p product) DeleteById(ctx *gofr.Context, id int) error {

	result, _ := ctx.DB().ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return nil
}
