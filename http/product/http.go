package product

import (
	"net/http"
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

func (h Handler) GetAllUsers(ctx *gofr.Context) (interface{}, error) {
	var prds []*models.Product
	products, err := h.Service.GetAllUsers(ctx)
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
