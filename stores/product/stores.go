package product

import (
	"database/sql"
	"gofr-curd/models"
	"gofr-curd/stores"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
}

func New() stores.Istore {
	return product{}
}

func (p product) GetProductByID(ctx *gofr.Context, id int) (*models.Product, error) {
	var prd models.Product

	rows := ctx.DB().QueryRowContext(ctx, "select * from Product where id = ?", id)

	err := rows.Scan(&prd.Id, &prd.Name, &prd.Type)

	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return &prd, nil
}

func (p product) GetAllProducts(ctx *gofr.Context) ([]*models.Product, error) {
	var prds []*models.Product

	rows, err := ctx.DB().QueryContext(ctx, "select * from Product")

	if err != nil {
		return []*models.Product{}, errors.Error("Couldn't execute query")
	}

	for rows.Next() {
		var prd models.Product
		_ = rows.Scan(&prd.Id, &prd.Name, &prd.Type)
		prds = append(prds, &prd)
	}

	return prds, nil
}

func (p product) CreateProduct(ctx *gofr.Context, prd models.Product) (int, error) {
	result, err := ctx.DB().ExecContext(ctx, "insert into Product(name,type) values (?,?)", prd.Name, prd.Type)
	if err != nil {
		return 0, errors.Error("Couldn't execute query")
	}

	newID, _ := result.LastInsertId()

	return int(newID), nil
}

func (p product) DeleteByID(ctx *gofr.Context, id int) error {
	res, err := ctx.DB().ExecContext(ctx, "delete from Product where id = ?", id)
	if err != nil {
		return errors.Error("Couldn't execute query")
	}

	r, _ := res.RowsAffected()

	if r == 0 {
		return errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id)}
	}

	return nil
}

func (p product) UpdateByID(ctx *gofr.Context, id int, prd models.Product) (int, error) {
	var i int

	fields, args := formUpdateQuery(prd)

	if fields == "" {
		return i, errors.Error("Nothing to Update")
	}

	fields1 := fields[:len(fields)-1]
	query1 := "update Product set" + fields1 + " where id = ?"

	args = append(args, id)

	res, err := ctx.DB().ExecContext(ctx, query1, args...)
	if err != nil {
		return i, errors.Error("Couldn't execute query")
	}

	r, _ := res.RowsAffected()

	if r == 0 {
		return i, errors.Error("SAME DATA GIVEN TO PREVIOUS DATA")
	}

	i = id

	return i, nil
}
