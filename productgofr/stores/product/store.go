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


func (p *DbStore) GetAllProduct(ctx *gofr.Context) ([]*models.Product,error) {
	var prods []*models.Product

	rows, err := ctx.DB().QueryContext(ctx,"SELECT * FROM product;")

	if err!=nil {
		return []*models.Product{},errors.Error("error")
	}
	
	for rows.Next() {
		var prod *models.Product

		err := rows.Scan(&prod.Id, &prod.Name,&prod.Type)

		if err!=nil {
			return nil, err
		}
		prods = append(prods,prod)
	}

	return prods,nil

}


func (p *DbStore) UpdateProduct(ctx *gofr.Context,pro models.Product) (*models.Product,error) {
	query := "UPDATE product SET"
	fields, values := formQuery(pro)
	query+= fields+ " WHERE id = ?"
	_,err := ctx.DB().Exec(query,values...)

	if err!=nil {
		return &models.Product{},errors.DB{Err:err}
	}

	return &pro,err
}


func (p *DbStore) DeleteProduct(ctx *gofr.Context,id int) error {
	_,err := ctx.DB().ExecContext(ctx,"DELETE FROM product where id = ?",id)
	if err!=nil {
		return errors.DB{Err: err}
	}
	return nil
}


func (p *DbStore) CreateProduct(ctx *gofr.Context,prod *models.Product) (*models.Product,error) {
	err := ctx.DB().QueryRowContext(ctx,"INSERT INTO product(Id, Name, Type) VALUES INTO(?,?,?)").Scan(&prod.Id,&prod.Name,&prod.Type)

	if err!=nil {
		return &models.Product{},errors.DB{Err:err}
	}

	return prod,nil
}