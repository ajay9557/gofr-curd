package stores

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetId(ctx *gofr.Context, id int) (models.Product, error)
}
