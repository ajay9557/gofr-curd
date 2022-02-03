package Stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

type Storeint interface {
	GetById(Id int, ctx *gofr.Context) (*model.Product, error)
}
