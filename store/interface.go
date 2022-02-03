package store

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetById(id int, ctx *gofr.Context) (models.Product, error)
}
