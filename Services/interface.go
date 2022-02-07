package Services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Serviceint interface {
	GetId(ctx *gofr.Context, Id int) (*model.Product, error)
	Create(ctx *gofr.Context, prod model.Product) (*model.Product, error)
	Delete(Id int, ctx *gofr.Context) error
	GetUser(ctx *gofr.Context) ([]*model.Product, error)
	Update(prod model.Product, ctx *gofr.Context) (*model.Product,error)
}
