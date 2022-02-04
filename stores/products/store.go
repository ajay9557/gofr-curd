package products

import (
	"database/sql"

	"errors"
	"strconv"

	perror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/arohanzst/testapp/models"
	mProduct "github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/stores"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct{}

func New() stores.Product {
	return &Product{}
}

//Fetching a Product using its id
func (p Product) ReadByID(ctx *gofr.Context, id int) (*mProduct.Product, error) {

	var resp models.Product

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM Product where Id = ?", id).Scan(&resp.Id, &resp.Name, &resp.Type)
	if err == sql.ErrNoRows {
		ctx.Log("ERRO", err)
		return nil, perror.EntityNotFound{Entity: "Product", ID: strconv.Itoa(id)}
	}

	return &resp, nil

}

//Fetching all the products
func (p Product) Read(ctx *gofr.Context) ([]models.Product, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT Id, Name, Type FROM Product")
	if rows == nil || err != nil {
		return nil, perror.DB{Err: err}
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err() // or modify return value
	}()

	products := make([]models.Product, 0)

	for rows.Next() {
		var c models.Product

		err := rows.Scan(&c.Id, &c.Name, &c.Type)
		if err != nil {
			return nil, err
		}

		products = append(products, c)
	}

	return products, nil
}

func (p Product) Create(ctx *gofr.Context, value *models.Product) (*models.Product, error) {

	row, err := ctx.DB().Exec("INSERT INTO Product(Name, Type) values(?, ?)", value.Name, value.Type)

	if err != nil {
		return nil, perror.DB{Err: errors.New("Error in given query")}
	}

	lastInsertId_64, err := row.LastInsertId()

	if err != nil {

		return nil, perror.DB{Err: errors.New("Error in given query")}
	}

	lastInsertId := int(lastInsertId_64)

	resp, err := p.ReadByID(ctx, lastInsertId)

	if err != nil {

		return nil, errors.New("Error in the given query")
	}
	return resp, nil
}

func (p Product) Update(ctx *gofr.Context, value *models.Product, id int) (*models.Product, error) {

	query := "Update Product Set "

	var arg []interface{}

	if value.Name != "" {

		query = query + "Name = ?,"
		arg = append(arg, value.Name)
	}

	if value.Type != "" {

		query = query + "Type = ?,"
		arg = append(arg, value.Type)

	}

	query = query[:len(query)-1]
	query = query + " where Id = ?"

	arg = append(arg, id)

	_, err := ctx.DB().ExecContext(ctx, query, arg...)
	if err != nil {
		return nil, perror.DB{Err: err}
	}

	resp, err := p.ReadByID(ctx, id)
	if err != nil {
		return nil, perror.DB{Err: err}
	}

	return resp, nil
}

func (p Product) Delete(ctx *gofr.Context, id int) error {

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM Product where Id=?", id)
	if err != nil {
		return perror.DB{Err: err}
	}

	return nil
}
