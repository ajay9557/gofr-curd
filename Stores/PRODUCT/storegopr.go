package PRODUCT

import (
	"errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Stores"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Datastore struct {
}

func New() Stores.Storeint {
	return &Datastore{}
}

func (store *Datastore) GetAllUser(ctx *gofr.Context) ([]*model.Product, error) {
	var prod []*model.Product

	query := "select id,name,type from user"

	rows, err := ctx.DB().QueryContext(ctx, query)

	if err != nil {
		return []*model.Product{}, errors.New("error in fetching the users")
	}

	for rows.Next() {
		var c model.Product
		err := rows.Scan(&c.Id, &c.Name, &c.Type)
		if err != nil {
			return nil, err
		}
		prod = append(prod, &c)
	}
	return prod, nil

}

func (store *Datastore) Create(prod model.Product, ctx *gofr.Context) (*model.Product, error) {
	var resp *model.Product
	query := "insert  into user (Id,Name,Type)  values(?,?,?)"
	row, err := ctx.DB().Exec(query, prod.Id, prod.Name, prod.Type) //.Scan(&resp.Id, &resp.Name, &resp.Type)
	if err != nil {
		return &model.Product{
			Id:   0,
			Name: "",
			Type: "",
		}, err
	}
	lid_64, _ := row.LastInsertId()
	lid := int(lid_64)
	resp, err = store.GetById(lid, ctx)
	if err != nil {
		return &model.Product{
			Id:   0,
			Name: "",
			Type: "",
		}, err
	}

	return resp, nil
}

func (store *Datastore) Update(prod model.Product, ctx *gofr.Context) (*model.Product, error) {
	feilds, val := QueryMod(prod)
	query := "UPDATE user SET " + feilds + "where id = ?"
	_, err := ctx.DB().Exec(query, val...)
	if err != nil {
		return nil, errors.New("error , no id provided , cannot update")
	}
	return store.GetById(prod.Id, ctx)
}

func (store *Datastore) Delete(id int, ctx *gofr.Context) error {
	query := "delete from user where id = ?"
	_, err := ctx.DB().Exec(query, id)
	if err != nil {
		return errors.New("error , cant delete ")
	}
	return nil

}

func (store *Datastore) GetById(Id int, ctx *gofr.Context) (*model.Product, error) {
	var solo *model.Product
	//err := ctx.DB().QueryContext(ctx,"SELECT Name,Type FROM user WHERE Id=?", Id).Scan(&solo.Id, &solo.Name, &solo.Type)
	err := ctx.DB().QueryRowContext(ctx, "SELECT Name , Type FROM user WHERE Id = ?", Id).Scan(&solo.Id)
	if err != nil {
		return nil, err
	}
	return solo, nil
}
