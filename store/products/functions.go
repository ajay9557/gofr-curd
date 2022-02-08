package products

import "github.com/ridhdhish-desai-zs/product-gofr/models"

func formUpdateQuery(p models.Product) (fields string, args []interface{}) {
	if p.Name != "" {
		fields += " name = ?,"

		args = append(args, p.Name)
	}

	if p.Category != "" {
		fields += " category = ?,"

		args = append(args, p.Category)
	}

	return
}
