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
	product, err := srv.storeInterface.GetProductById(ctx, convId)
	if err != nil {
		return prd, err
	}
	prd = product
	return prd, nil
}

func (srv *ProductService) GetAllProducts(ctx *gofr.Context) ([]*models.Product, error) {
	var prd []*models.Product
	res, err := srv.storeInterface.GetAllProducts(ctx)
	if err != nil {
		return prd, err
	}
	prd = res
	return prd, nil
}

func (srv *ProductService) CreateProduct(ctx *gofr.Context, prd models.Product) (*models.Product, error) {
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
	updatedUser, _ := srv.storeInterface.GetProductById(ctx, id)
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

	// var prd1 *models.Product

	// if reflect.DeepEqual(models.Product{}, prd) {
	// 	return prd1, errors.Error("Given Empty data")
	// }
	// if prd.Name == "" {
	// 	return prd1, errors.Error("Please provide Data for Name")
	// }

	// if prd.Type == "" {
	// 	return prd1, errors.Error("Please provide Data for Type")
	// }

	_, err = srv.storeInterface.UpdateById(ctx, convId, prd)
	if err != nil {
		return prd1, err
	}
	updatedUser, _ := srv.storeInterface.GetProductById(ctx, convId)
	prd1 = updatedUser
	return prd1, nil

}

// Not checking whether the an entity exists for a given Id
// It will be directly taken care in store layer
