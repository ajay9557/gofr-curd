package stores

import (
	"productGofr/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Istore interface {
	UserById(ctx *gofr.Context, id int) (*models.Product, error)
	GetAllUsers(ctx *gofr.Context) ([]*models.Product, error)
}
