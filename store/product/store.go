package product

import (
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

	if err != nil {
		return product, errors.Error("product data not found for the given id")
	}

	return product, nil
}

func (p dbStore) GetAllProducts(ctx *gofr.Context) ([]models.Product, error) {
	var prod []models.Product

	query := "select id, name, type from product"

	rows, err := ctx.DB().QueryContext(ctx, query)

	if err != nil {
		return nil, errors.Error("internal db error")
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		_ = rows.Scan(&product.ID, &product.Name, &product.Type)

		prod = append(prod, product)
	}

	return prod, nil
}

func (p dbStore) UpdateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "update product set name = ?, type = ? where id = ?", prod.Name, prod.Type, prod.ID)

	if err != nil {
		return models.Product{}, errors.Error("error while updating product")
	}

	return prod, nil
}

func (p dbStore) CreateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "insert into product values (?,?,?)", prod.ID, prod.Name, prod.Type)

	if err != nil {
		return prod, errors.Error("internal db error")
	}

	return prod, nil
}

func (p dbStore) DeleteProduct(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from product where id = ?", id)

	if err != nil {
		return errors.Error("internal db error")
	}

	return nil
}
