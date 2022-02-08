package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/arohanzst/testapp/models"
)

type Product interface {
	ReadByID(ctx *gofr.Context, id int) (*models.Product, error)
	Read(ctx *gofr.Context) ([]models.Product, error)
	Create(ctx *gofr.Context, value *models.Product) (*models.Product, error)
	Update(ctx *gofr.Context, value *models.Product, id int) (*models.Product, error)
	Delete(ctx *gofr.Context, id int) error
}
