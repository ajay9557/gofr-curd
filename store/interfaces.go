package store

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetById(ctx *gofr.Context, id int) (*models.Product, error)
}
