package product

import (
	"fmt"
	models "zopsmart/productgofr/models"
)

func formQuery(prod models.Product) (string, []interface{}) {
	var query string
	var values []interface{}

	if prod.Name=="" {
		query+="Name = ?,"
		values = append(values, prod.Name)
	}

	if prod.Type=="" {
		query+="Type = ?,"
		values = append(values, prod.Type)

	}

	query = query[:len(query)-1]
	fmt.Println(query)
	values = append(values, prod.Id)

	return query,values

}