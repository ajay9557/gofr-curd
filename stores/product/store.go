package product

import (
	"database/sql"
	"gofr-curd/model"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type productstorer struct {
}

func New() productstorer {
	return productstorer{}
}

func (p productstorer) GetProductById(id int, ctx *gofr.Context) (model.ProductDetails, error) {

	prod := model.ProductDetails{}
	err := ctx.DB().QueryRowContext(ctx, "select Id,Name,Types from Product where Id=?", id).Scan(&prod.Id, &prod.Name, &prod.Types)
	if err == sql.ErrNoRows {
		return model.ProductDetails{}, errors.EntityNotFound{Entity: "Product", ID: strconv.Itoa(id)}
	}
	return prod, nil
}

func (p *productstorer) DeleteProductId(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "Delete from Product where Id=?", id)
	if err != nil {
		return errors.Error("Internal Database error")
	}
	return nil
}
func (p *productstorer) UpdateProductById(ctx *gofr.Context, prod model.ProductDetails) error {
	_, err := ctx.DB().ExecContext(ctx, "Update Product set Name=?,Types=? where Id=?", prod.Name, prod.Types, prod.Id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}

func (p *productstorer) CreateProducts(ctx *gofr.Context, product model.ProductDetails) (model.ProductDetails, error) {
	_, err := ctx.DB().ExecContext(ctx, "Insert into Product values(?,?,?)", product.Id, product.Name, product.Types)
	if err != nil {
		return product, errors.DB{Err: err}
	}
	return product, nil
}

func (p *productstorer) GetAll(ctx *gofr.Context) ([]model.ProductDetails, error) {
	var products []model.ProductDetails
	rows, err := ctx.DB().QueryContext(ctx, "Select Id,Name,Types from Product")
	if err != nil {
		return nil, errors.DB{}
	}
	defer rows.Close()
	for rows.Next() {
		var product model.ProductDetails
		err := rows.Scan(&product.Id, &product.Name, &product.Types)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
