package product

import (
	//	"database/sql"
	"strconv"
	models "zopsmart/productgofr/models"
	stores "zopsmart/productgofr/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type DBstore struct{}

func New() stores.Store {
	return &DBstore{}
}

func (p *DBstore) GetProdByID(ctx *gofr.Context, id int) (models.Product, error) {
	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, "Select id,name,type from product where id =?", id).Scan(&product.Id, &product.Name, &product.Type)

	if err != nil {
		return product, errors.EntityNotFound{Entity: "product", ID: strconv.Itoa(id)}
	}
	return product, nil
}

func (p *DBstore) DeleteProduct(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "Delete from product where id=?", id)
	if err != nil {
		return errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}
	return nil
}

func (p *DBstore) UpdateProduct(ctx *gofr.Context, prod models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "Update product set name=?,type=? where id=?", prod.Name, prod.Type, prod.Id)
	if err != nil {
		return errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(prod.Id)}
	}
	return nil
}

func (p *DBstore) CreateProduct(ctx *gofr.Context, product models.Product) error {
	_, err := ctx.DB().ExecContext(ctx, "insert into product values(?,?,?)", product.Id, product.Name, product.Type)
	if err != nil {
		return errors.EntityNotFound{
			Entity: "product",
			ID:     "id",
		}
	}
	return nil
}

func (p *DBstore) GetAllProduct(ctx *gofr.Context) ([]models.Product, error) {
	var products []models.Product
	rows, err := ctx.DB().QueryContext(ctx, "Select id,name,type from product")
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "products", ID: "all"}
	}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Type)
		if err != nil {
			return nil, errors.Error("internal DB error")
		}
		products = append(products, product)
	}
	return products, nil
}
