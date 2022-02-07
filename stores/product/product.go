package product

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	goError "errors"
	"product/models"
)

type Store struct {
}

func New() Store {
	return Store{}
}

func (store *Store) GetProductById(ctx *gofr.Context, id int) (models.Product, error) {
	product := models.Product{}
	_ = ctx.DB().QueryRow("select * from product where id=?", id).Scan(&product.Id, &product.Name, &product.Type)
	return product, nil
}

func (store *Store) GetAllProduct(ctx *gofr.Context) ([]models.Product, error) {
	productList := []models.Product{}
	row, _ := ctx.DB().QueryContext(ctx, "select * from product")

	defer func() {
		_ = row.Close()
	}()

	for row.Next() {
		product := models.Product{}
		_ = row.Scan(&product.Id, &product.Name, &product.Type)
		productList = append(productList, product)
	}
	return productList, nil
}

func (store *Store) AddProduct(ctx *gofr.Context, product models.Product) error {
	_, err := ctx.DB().Exec("insert into product(name, type) values(?, ?)", product.Name, product.Type)
	if err != nil || product.Name == "" || product.Type == "" {
		return goError.New("FAILED TO ADD PRODUCT")
	}
	return nil
}

func (store *Store) UpdateProduct(ctx *gofr.Context, product models.Product) error {
	query := "update product set"
	var args []interface{}

	if product.Name != "" {
		query += " name=?,"
		args = append(args, product.Name)
	}

	if product.Type != "" {
		query += " type=?"
		args = append(args, product.Type)
	}

	if product.Id > 0 {
		query += " where id=?"
		args = append(args, product.Id)
	}
	_, err := ctx.DB().Exec(query, args...)
	if err != nil {
		return goError.New("FAILED TO UPDATE THE PRODUCT")
	}
	return nil
}

func (store *Store) DeleteProduct(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from product where id=?", id)
	if err != nil || id < 0 {
		return goError.New("FAILED TO DELETE PRODUCT")
	}
	return nil
}
