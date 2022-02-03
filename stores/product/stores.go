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
