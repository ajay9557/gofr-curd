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

func (srv *ProductService) CreateProduct(ctx *gofr.Context, prd models.Product) (*models.Product, error) {
	var prd1 *models.Product
	id, err := srv.storeInterface.CreateProduct(ctx, prd)
	if err != nil {
		return prd1, err
	}
	updatedUser, _ := srv.storeInterface.UserById(ctx, id)
	prd1 = updatedUser
	return prd1, nil

}

func (srv *ProductService) DeleteById(ctx *gofr.Context, id string) error {
	// var prd *models.Product
	convId, err := strconv.Atoi(id)
	if err != nil {
		return errors.MissingParam{Param: []string{id}}
	}
	if convId < 0 {
		return errors.InvalidParam{Param: []string{id}}

		// return &prd, errors.EntityNotFound{Entity: "products", ID: "id"}
	}
	err = srv.storeInterface.DeleteById(ctx, convId)
	if err != nil {
		return err
	}
	// prd = product
	return nil
}

func (srv *ProductService) UpdateById(ctx *gofr.Context, id string, prd models.Product) (*models.Product, error) {
	var prd1 *models.Product
	convId, err := strconv.Atoi(id)
	if err != nil {
		return prd1, errors.MissingParam{Param: []string{id}}
	}
	if convId < 0 {
		return prd1, errors.InvalidParam{Param: []string{id}}

		// return &prd, errors.EntityNotFound{Entity: "products", ID: "id"}
	}

	_, err = srv.storeInterface.UpdateById(ctx, convId, prd)
	if err != nil {
		return prd1, err
	}
	updatedUser, _ := srv.storeInterface.UserById(ctx, convId)
	prd1 = updatedUser
	return prd1, nil

}
