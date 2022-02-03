package services

import (
	"productGofr/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Iservice interface {
	GetProductById(ctx *gofr.Context, id string) (*models.Product, error)
	GetAllUsers(ctx *gofr.Context) ([]*models.Product, error)
}
