package services

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByUserId(ctx *gofr.Context, id int) (models.Product, error)
}
