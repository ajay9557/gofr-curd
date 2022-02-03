package product

import (
	"database/sql"
	"productGofr/models"
	"productGofr/stores"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
}

func New() stores.Istore {
	return product{}
}

func (p product) UserById(ctx *gofr.Context, id int) (*models.Product, error) {
	var prd models.Product
	rows := ctx.DB().QueryRow("select * from Product where id = ?", id)
	// if rows.Err() != nil {
	// 	return nil, rows.Err()
	// }
	err := rows.Scan(&prd.Id, &prd.Name, &prd.Type)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(id) /*fmt.Sprint(id)*/}
	}
	return &prd, nil
}

func (p product) GetAllUsers(ctx *gofr.Context) ([]*models.Product, error) {
	var prds []*models.Product
	rows, _ := ctx.DB().Query("select * from Product")
	// if err != nil {
	// 	return []*models.Product{}, errors.DB{Err: err}
	// }
	for rows.Next() {
		var prd models.Product
		err := rows.Scan(&prd.Id, &prd.Name, &prd.Type)
		if err != nil {
			return prds, errors.EntityNotFound{Entity: "Product"}
		}

		prds = append(prds, &prd)
	}

	return prds, nil

}

func (p product) CreateProduct(ctx *gofr.Context, prd models.Product) (int, error) {
	result, err := ctx.DB().Exec("insert into Product(name,type) values (?,?)", prd.Name, prd.Type)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	newId, _ := result.LastInsertId()

	return int(newId), nil

}

func (p product) DeleteById(ctx *gofr.Context, id int) error {
	// var prd models.Product
	_, err := ctx.DB().Exec("delete from Product where id = ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}

func (p product) UpdateById(ctx *gofr.Context, id int, prd models.Product) (int, error) {
	var i int

	_, err := ctx.DB().Exec("update Product set name = ?,type = ? where id = ?", prd.Name, prd.Type, id)
	if err != nil {
		return i, errors.DB{Err: err}
	}
	i = id

	return i, nil
}
