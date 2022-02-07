package product 
import (
	models "zopsmart/productgofr/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)
type Store interface {
	GetProdByID(ctx *gofr.Context, id int) (*models.Product, error)
	DeleteProduct(ctx *gofr.Context,id int) error
	CreateProduct(ctx *gofr.Context,prod *models.Product) (*models.Product,error)
	UpdateProduct(ctx *gofr.Context,pro models.Product) (*models.Product,error)
	GetAllProduct(ctx *gofr.Context) ([]*models.Product,error)
}