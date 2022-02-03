package product

import (
	"productGofr/models"
	"productGofr/services"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	Service services.Iservice
}

func (h Handler) GetByIdHandler(ctx *gofr.Context) (interface{}, error) {
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
		return nil, err
	}
	resData := struct {
		Product *models.Product `json : "product"`
	}{
		Product: p,
	}

	return resData, nil
}
