package PRODUCT

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Stores"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Datastore struct{}

func New() Stores.Storeint {
	return Datastore{}
}

func (store Datastore) GetById(Id int, ctx *gofr.Context) (*model.Product, error) {
	var solo model.Product
	err := ctx.DB().QueryRow("SELECT Id,Name,Type FROM user WHERE Id=?", Id).Scan(&solo.Id, &solo.Name, &solo.Type)
	if err != nil {
		return nil, err
	}
	return &solo, nil
}
