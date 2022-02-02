package product

import (
	"fmt"
	"product/models"
	"product/services"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type HttpService struct {
	Service services.Service
}

func (service *HttpService) GetByIdHandler(ctx *gofr.Context) (interface{}, error) {
	fmt.Println("called")
	productId := ctx.PathParam("id")
	var product models.Product
	// if strings.TrimSpace(productId) == "" {
	// 	// errorResponse := models.ErrorResponse{
	// 	// 	Code:    http.StatusBadRequest,
	// 	// 	Message: "INVALID ID",
	// 	// }
	// 	return nil, errors.EntityNotFound{Entity: "product", }
	// }

	id, err := strconv.Atoi(productId)
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "product", ID: productId}
	}

	product, _ = service.Service.GetProductById(ctx, id)

	// responseObj := models.Response{
	// 	Data:       product,
	// 	Message:    "Product Found",
	// 	StatusCode: http.StatusOK,
	// }
	return struct {
		Product models.Product `json:"product"`
	}{
		Product: product,
	}, nil
}
