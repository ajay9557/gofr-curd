package product

import (
	"strings"

	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

func buildQuery(product *models.Product) (feilds string, values []interface{}) {
	if product.Name != "" {
		feilds += "name = ?, "

		values = append(values, product.Name)
	}

	if product.Type != "" {
		feilds += "type = ?, "

		values = append(values, product.Type)
	}

	feilds = strings.TrimRight(feilds, ", ")

	values = append(values, product.ID)

	return feilds, values
}
