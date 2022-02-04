package product

import (
	// "encoding/json"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/services"
	"net/http"

	// "reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Service services.Iservice
}

func (h Handler) GetProductByIdHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")
	// if param == ""{
	// 	return nil, errors.MissingParam{Param: []string{"id"}}
	// }
	// id,err := strconv.Atoi(param)
	// if err != nil{
	// 	return nil,errors.InvalidParam{Param: []string{"id"}}
	// }
	p, err := h.Service.GetProductById(ctx, param)
	if err != nil {
		return models.Response{}, err
	}
	// resData := struct {
	// 	Product *models.Product `json : "product"`
	// }{
	// 	Product: p,
	// }

	responseObj := models.Response{
		Data:       &p,
		Message:    "Successfully Product Retrived",
		StatusCode: http.StatusOK,
	}

	// return resData, nil
	return responseObj, nil
}

func (h Handler) GetAllProductsHandler(ctx *gofr.Context) (interface{}, error) {
	var prds []*models.Product
	products, err := h.Service.GetAllProducts(ctx)
	if err != nil {
		return prds, err
	}
	prds = products
	responseObj := models.Response{
		Data:       &prds,
		Message:    "Successfully Product Retrived",
		StatusCode: http.StatusOK,
	}
	return responseObj, nil
}

func (h Handler) CreateProductHandler(ctx *gofr.Context) (interface{}, error) {
	var prd models.Product

	// err := json.NewDecoder(ctx.Request().Body).Decode(&prd)
	// if err != nil || reflect.DeepEqual(prd, models.Product{}) {
	// 	return prd, err
	// }

	if err := ctx.Bind(&prd); err != nil {
		fmt.Print("Anusri")
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}

	}

	NewPrd, err := h.Service.CreateProduct(ctx, prd)
	if err != nil {
		return models.Response{}, err
	}
	responseObj := models.Response{
		Data:       &NewPrd,
		Message:    "Successfully Product Created",
		StatusCode: http.StatusOK,
	}

	// return resData, nil
	return responseObj, nil

}

func (h Handler) DeleteByIdHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")
	err := h.Service.DeleteById(ctx, param)
	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Message:    "Successfully Product Deleted",
		StatusCode: http.StatusOK,
	}

	// return resData, nil
	return responseObj, nil
}

func (h Handler) UpdateByIdHandler(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")

	var prd models.Product
	if err := ctx.Bind(&prd); err != nil {
		fmt.Print("Anusri")
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}

	}

	p, err := h.Service.UpdateById(ctx, param, prd)
	if err != nil {
		return models.Response{}, err
	}

	responseObj := models.Response{
		Data:       &p,
		Message:    "Successfully Product Updated",
		StatusCode: http.StatusOK,
	}

	// return resData, nil
	return responseObj, nil

}
