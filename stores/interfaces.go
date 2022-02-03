package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/arohanzst/testapp/models"
)

type Product interface {
	ReadByID(ctx *gofr.Context, id int) (*models.Product, error)
}
