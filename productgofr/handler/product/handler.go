package product

import (
	"strconv"

	services "zopsmart/productgofr/services"
	//stores "zopsmart/productgofr/stores"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProdHandler struct {
	serv services.Services
}

func New(s services.Services) ProdHandler{
	return ProdHandler{serv:s}
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