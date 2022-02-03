package Services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Serviceint interface {
	GetId(ctx *gofr.Context, Id int) (*model.Product, error)
}
