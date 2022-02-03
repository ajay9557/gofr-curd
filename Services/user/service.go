package user

import (
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
