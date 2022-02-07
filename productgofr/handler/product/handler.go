package product

import (
	"strconv"
	"io/ioutil"
	services "zopsmart/productgofr/services"
	models "zopsmart/productgofr/models"
	"encoding/json"
	//stores "zopsmart/productgofr/stores"
//	"developer.zopsmart.com/go/gofr/examples/using-cassandra/models"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProdHandler struct {
	serv services.Services
}

func New(s services.Services) ProdHandler{
	return ProdHandler{serv:s}
}

func(h ProdHandler) GetAllProductHandler(ctx *gofr.Context) (interface{},error) {
	 res,err := h.serv.GetAllProd(ctx)

	 if err!=nil {
		 return nil,err
	}
	
	return res,nil

}

func(h ProdHandler) GetProdByIdHandler(ctx * gofr.Context) (interface{},error){
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.serv.GetProdByID(ctx, id)
	if err != nil {
		return nil,err
	}

	return resp, nil
}

func (h ProdHandler) CreateProdHandler(ctx *gofr.Context) (interface{},error) {
	var pro *models.Product

	if err:= ctx.Bind(&pro); err!=nil {
		ctx.Logger.Errorf("error in binding: %v",err)
		return nil,errors.InvalidParam{Param: []string{"body"}}
	}

	if pro.Id!=0 {
		return nil,errors.InvalidParam{Param: []string{"id"}}
	}
	res,err := h.serv.CreateProduct(ctx,pro)
	if err!=nil {
		return nil,err
	}

	return res,nil
}

func (h ProdHandler) DeleteProductHandler(ctx *gofr.Context) (interface{},error) {
	i := ctx.PathParam("id")

	if i =="" {
		return nil,errors.MissingParam{Param: []string{"id"}}
	}

	id,err := strconv.Atoi(i)
	if err!=nil {
		return nil,errors.InvalidParam{Param: []string{"id"}}
	}
	if err:=h.serv.DeleteProduct(ctx,id); err!=nil {
		return nil,err
	}

	return "Deleted Successfully",nil

}

func (h ProdHandler) UpdateProduct(ctx *gofr.Context) (interface{},error) {
	var prod models.Product
	resBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil,err
	}
	err = json.Unmarshal(resBody, &prod)
	if err != nil {
		return nil,err
	}
	if prod.Id == 0 {
		return nil,err
	}
	res, err:=h.serv.UpdateProduct(ctx,prod);
	if err!=nil {
		return nil,err
	}

	return res,nil

}