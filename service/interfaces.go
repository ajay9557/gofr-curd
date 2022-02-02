package service

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Services interface {
	GetById(ctx *gofr.Context, id int) (*models.Product, error)
}
