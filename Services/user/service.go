package user

import (
	"errors"
	"log"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Stores"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Product struct {
	u Stores.Storeint
}

func New(pro Stores.Storeint) Product {
	return Product{
		u: pro,
	}
}

func (u Product) GetId(ctx *gofr.Context, Id int) (*model.Product, error) {
	if Id <= 0 {
		log.Panicln("Invalid Id")
	}
	prod, err := u.u.GetById(Id, ctx)

	if err != nil {
		return prod, err
	}
	return prod, nil
}

func (u Product) Create(ctx *gofr.Context, prod model.Product) (*model.Product, error) {
	if prod.Id <= 0 {
		return nil, errors.New("error invalid id")
	}

	if prod.Name == "" {
		return nil, errors.New("invalid name ")
	}

	if prod.Type == "" {
		return nil, errors.New("invalid type")
	}

	prods, err := u.u.Create(prod, ctx)
	if err != nil {
		return nil, errors.New("fail to execute")
	}
	return prods, nil

}

func (u Product) Delete(Id int, ctx *gofr.Context) error {
	if Id <= 0 {
		return errors.New("error invalid id")
	}
	return u.u.Delete(Id, ctx)
}

func (u Product) GetUser(ctx *gofr.Context) ([]*model.Product, error) {

	prods, err := u.u.GetAllUser(ctx)
	if err != nil {
		return nil, errors.New("fail to execute")
	}
	return prods, nil
}

func (u Product) Update(prod model.Product, ctx *gofr.Context) (*model.Product, error) {
	if prod.Id <= 0 {
		return nil, errors.New("error invalid id")
	}
	if prod.Name == "" {
		return nil, errors.New("invalid name ")
	}
	if prod.Type == "" {
		return nil, errors.New("invalid type")
	}
	return u.u.Update(prod, ctx)
}
