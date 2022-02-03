package product

import (
	"database/sql"
	"fmt"
	models "zopsmart/productgofr/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/aws/aws-sdk-go/private/protocol/query"
)

type DbStore struct {
	db *sql.DB
}

func New() *DbStore {
	return &DbStore{
	}
}

func (s *DbStore) GetProdByID(ctx *gofr.Context, id int) (*models.Product, error) {
	var resp models.Product

	err := ctx.DB().QueryRowContext(ctx, " SELECT * FROM product where id=?", id).Scan(&resp.Id, &resp.Name, &resp.Type)

	if err != nil {
		return &models.Product{}, errors.EntityNotFound{Entity: "Product", ID: fmt.Sprint(id)}
	}

	return &resp, nil

}
