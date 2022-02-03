package product

import (
	"productGofr/models"
	"productGofr/services"
	"productGofr/stores"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductService struct {
	storeInterface stores.Istore
}

func New(si stores.Istore) services.Iservice {
	return &ProductService{storeInterface: si}
}

func (srv *ProductService) GetProductById(ctx *gofr.Context, id string) (*models.Product, error) {
	var prd *models.Product
	convId, err := strconv.Atoi(id)
	if err != nil {
		return prd, errors.MissingParam{Param: []string{id}}
	}
	if convId < 0 {
		return prd, errors.InvalidParam{Param: []string{id}}

		// return &prd, errors.EntityNotFound{Entity: "products", ID: "id"}
	}
	product, err := srv.storeInterface.UserById(ctx, convId)
	if err != nil {
		return prd, err
	}
	prd = product
	return prd, nil

}

func (srv *ProductService) GetAllUsers(ctx *gofr.Context) ([]*models.Product, error) {
	var prd []*models.Product
	res, err := srv.storeInterface.GetAllUsers(ctx)
	if err != nil {
		return prd, err
	}
	prd = res
	return prd, nil
}
