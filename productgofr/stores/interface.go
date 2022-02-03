package product 
import (
	models "zopsmart/productgofr/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)
type Store interface {
	GetProdByID(ctx *gofr.Context, id int) (*models.Product, error)
}