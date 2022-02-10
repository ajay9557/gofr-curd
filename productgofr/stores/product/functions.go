package product

import (
	models "zopsmart/productgofr/models"
)

func formQuery(prod models.Product) (query string,values []interface{}) {
		if prod.Name != "" {
		query += " name = ?,"
		values = append(values, prod.Name)
	}

	if prod.Type != "" {
		query += " category = ?,"
		values = append(values, prod.Type)
	}

	return

}