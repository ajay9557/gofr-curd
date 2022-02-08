package product

import (
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"strings"
)

func buildQuery(product *models.Product) (string, []interface{}) {
	feilds := ""
	var values []interface{}

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
