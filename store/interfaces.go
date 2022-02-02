package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"gofrPractice/models"
)

type Store interface {
	GetById(ctx *gofr.Context, id int) (*models.Product, error)
}
