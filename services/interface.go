package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

// go:todo mockgen -package=services --source=interface.go --destination=mock_interface.go
type Product interface {
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	//Update(ctx *gofr.Context, id int) (*models.Product, error)
	//Delete(ctx *gofr.Context, id int) (*models.Product, error)

	//GetAll(ctx *gofr.Context) ([]*models.Product,error)

}
