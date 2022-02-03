package products

import (
	"database/sql"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
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
		return nil, errors.EntityNotFound{Entity: "Product", ID: strconv.Itoa(id)}
	}

	return &resp, nil

}
