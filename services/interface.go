package services

import (
	"product/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetProductById(ctx *gofr.Context, id int) (models.Product, error)
}
