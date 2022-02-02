package stores

import (
	"product/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetProductById(ctx *gofr.Context, id int) (models.Product, error)
}
