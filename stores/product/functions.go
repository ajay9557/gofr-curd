package product

import (
	"gofr-curd/models"
)

// func formUpdateQuery(u models.Product) (string, []interface{}) {
// 	var where string = ""
// 	var value []interface{}
// 	if u.Name != "" {
// 		where = where + "name = ?,"
// 		value = append(value, u.Name)
// 	}
// 	if u.Type != "" {
// 		where = where + "type = ?,"
// 		value = append(value, u.Type)
// 	}
// 	return where, value
// }

func formUpdateQuery(p models.Product) (fields string, args []interface{}) {
	if p.Name != "" {
		fields += " name = ?,"

		args = append(args, p.Name)
	}

	if p.Type != "" {
		fields += " type = ?,"

		args = append(args, p.Type)
	}

	return
}
