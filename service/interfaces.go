package service

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"gofrPractice/models"
)

type Services interface {
	GetById(ctx *gofr.Context, id int) (*models.Product, error)
}
