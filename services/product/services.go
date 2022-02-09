package product

import (
	"gofr-curd/models"
	"gofr-curd/services"
	"gofr-curd/stores"
	"reflect"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service struct {
	storeInterface stores.Istore
}

func New(si stores.Istore) services.Iservice {
	return &Service{storeInterface: si}
}

func (srv *Service) GetProductByID(ctx *gofr.Context, id string) (*models.Product, error) {
	var prd *models.Product

	convID, err := strconv.Atoi(id)

	if err != nil {
		return prd, errors.MissingParam{Param: []string{id}}
	}

	if convID < 0 {
		return prd, errors.InvalidParam{Param: []string{id}}
	}

	product, err := srv.storeInterface.GetProductByID(ctx, convID)

	if err != nil {
		return prd, err
	}

	prd = product

	return prd, nil
}

func (srv *Service) GetAllProducts(ctx *gofr.Context) ([]*models.Product, error) {
	var prd []*models.Product

	res, err := srv.storeInterface.GetAllProducts(ctx)

	if err != nil {
		return prd, err
	}

	prd = res

	return prd, nil
}

func (srv *Service) CreateProduct(ctx *gofr.Context, prd models.Product) (*models.Product, error) {
	var prd1 *models.Product

	if reflect.DeepEqual(models.Product{}, prd) {
		return prd1, errors.Error("Given Empty data")
	}

	if prd.Name == "" {
		return prd1, errors.Error("Please provide Data for Name")
	}

	if prd.Type == "" {
		return prd1, errors.Error("Please provide Data for Type")
	}

	id, err := srv.storeInterface.CreateProduct(ctx, prd)
	if err != nil {
		return prd1, err
	}
	// Fetchinf the created product
	updatedUser, _ := srv.storeInterface.GetProductByID(ctx, id)
	prd1 = updatedUser

	return prd1, nil
}

func (srv *Service) DeleteByID(ctx *gofr.Context, id string) error {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return errors.MissingParam{Param: []string{id}}
	}

	if convID < 0 {
		return errors.InvalidParam{Param: []string{id}}
	}

	err = srv.storeInterface.DeleteByID(ctx, convID)
	if err != nil {
		return err
	}
	// prd = product
	return nil
}

func (srv *Service) UpdateByID(ctx *gofr.Context, id string, prd models.Product) (*models.Product, error) {
	var prd1 *models.Product

	convID, err := strconv.Atoi(id)

	if err != nil {
		return prd1, errors.MissingParam{Param: []string{id}}
	}

	if convID < 0 {
		return prd1, errors.InvalidParam{Param: []string{id}}
	}

	_, err = srv.storeInterface.UpdateByID(ctx, convID, prd)
	if err != nil {
		return prd1, err
	}

	updatedUser, _ := srv.storeInterface.GetProductByID(ctx, convID)
	prd1 = updatedUser

	return prd1, nil
}

// Not checking whether the an entity exists for a given Id
// It will be directly taken care in store layer
