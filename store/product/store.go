package product

import (
	"database/sql"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/store"
	"strings"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type productStorer struct {
}

func New() store.Store {
	return productStorer{}
}

func (p productStorer) GetById(id int, ctx *gofr.Context) (models.Product, error) {

	ReadQ := "Select Id,Name,Type from product where Id=?"
	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, ReadQ, id).Scan(&product.Id, &product.Name, &product.Type)
	if err == sql.ErrNoRows {
		return models.Product{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(id)}
	}
	return product, nil
}

func (p productStorer) GetAllProducts(ctx *gofr.Context) ([]models.Product, error) {
	var products []models.Product
	var product models.Product
	ReadQ := "Select Id,Name,Type from product;"
	rows, err := ctx.DB().QueryContext(ctx, ReadQ)
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&product.Id, &product.Name, &product.Type)
		products = append(products, product)
	}
	return products, nil
}

func (p productStorer) InsertProduct(product models.Product, ctx *gofr.Context) error {
	insertQ := "insert into product(Id,Name,Type) Values(?,?,?)"
	_, err := ctx.DB().ExecContext(ctx, insertQ, product.Id, product.Name, product.Type)
	if err != nil {
		return errors.Error("Error in executing query")
	}
	return nil
}

func (p productStorer) UpdateProduct(product models.Product, ctx *gofr.Context) error {
	updateEntities := []interface{}{}
	updateQ := "update product set " //Update query

	if product.Name != "" {
		updateQ += "Name=?,"
		updateEntities = append(updateEntities, product.Name)
	}
	if product.Type != "" {
		updateQ += "Type=?,"
		updateEntities = append(updateEntities, product.Type)
	}
	updateQ = strings.TrimRight(updateQ, ",")
	if product.Id != 0 {
		updateQ += " where Id=?;"
		updateEntities = append(updateEntities, product.Id)
	}
	_, err := ctx.DB().ExecContext(ctx, updateQ, updateEntities...)
	if err != nil {
		return errors.Error("Error in executing query")
	}
	return nil
}

func (p productStorer) DeleteById(id int, ctx *gofr.Context) error {
	deleteQ := "delete from product where Id=?"
	_, err := ctx.DB().ExecContext(ctx, deleteQ, id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
