package Stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Storeint interface {
	GetById(Id int, ctx *gofr.Context) (*model.Product, error)
	GetAllUser(ctx *gofr.Context) ([]*model.Product, error)
	Create(prod model.Product, ctx *gofr.Context) (*model.Product, error)
	Update(prod model.Product, ctx *gofr.Context) (*model.Product, error)
	Delete(id int, ctx *gofr.Context) error
}
