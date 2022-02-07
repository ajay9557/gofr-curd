package products

import (
	"database/sql"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductStorer struct {
}

func New() stores.Store {
	return &ProductStorer{}
}

func (p *ProductStorer) GetID(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product
	err := ctx.DB().QueryRowContext(ctx, "Select ID,Name,Type from Product where Id =?", id).Scan(&product.ID, &product.Name, &product.Type)

	if err == sql.ErrNoRows {
		return product, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
	}

	return product, nil
}

func (p *ProductStorer) DeleteID(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "Delete from Product where id=?", id)

	if err != nil {
		return errors.Error("Internal DB error")
	}

	return nil
}

func (p *ProductStorer) UpdateID(ctx *gofr.Context, product models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "Update Product set Name=?,Type=? where Id=?", product.Name, product.Type, product.ID)

	if err != nil {
		return errors.Error("Internal DB error")
	}

	return nil
}

func (p *ProductStorer) CreateProducts(ctx *gofr.Context, product models.Product) (models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "Insert into Product values(?,?,?)", product.ID, product.Name, product.Type)

	if err != nil {
		return product, errors.Error("Internal DB error")
	}

	return product, nil
}

func (p *ProductStorer) GetAll(ctx *gofr.Context) ([]models.Product, error) {
	var products []models.Product

	rows, err := ctx.DB().QueryContext(ctx, "Select Id,Name,Type from Product")

	if err != nil {
		return nil, errors.Error("Internal DB error")
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
