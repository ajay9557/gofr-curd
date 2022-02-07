package PRODUCT

import "github.com/shaurya-zopsmart/Gopr-devlopment/model"

func QueryMod(data model.Product) (string, []interface{}) {
	var q string
	var val []interface{}

	if data.Id < 0 {
		return "", nil
	}
	if data.Name != "" {
		q += " name = ?,"
		val = append(val, data.Name)
	}
	if data.Type != "" {
		q += " type = ?"
		val = append(val, data.Type)
	}

	val = append(val, data.Id)
	return q, val
}
