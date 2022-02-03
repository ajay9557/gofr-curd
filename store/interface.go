package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
)

type ProductStore interface {
	GetProductById(ctx *gofr.Context, id int) (models.Product, error)
}
