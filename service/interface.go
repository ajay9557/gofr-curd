package service

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByProductId(id int, ctx *gofr.Context) (models.Product, error)
}
