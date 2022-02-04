package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

type Store interface {
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	//Update(ctx *gofr.Context, id int) (*models.Product, error)
	//Delete(ctx *gofr.Context, id int) (*models.Product, error)
}
