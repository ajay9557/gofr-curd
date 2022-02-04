package products

import (
	"strconv"
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

func validateId(id string) (bool, error) {
	if id == "" {
		return false, errors.MissingParam{Param: []string{"id"}}
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return false, errors.InvalidParam{Param: []string{"id"}}
	}
	if i < 1 {
		return false, errors.InvalidParam{Param: []string{"id"}}
	}

	return true, nil
}

func CheckBody(prod model.Product) (bool, error) {
	if prod.Id != 0 {
		return false, errors.InvalidParam{Param: []string{"id"}}
	}
	if prod.Name == "" {
		return false, errors.MissingParam{Param: []string{"Name"}}
	}
	if prod.Type == "" {
		return false, errors.MissingParam{Param: []string{"Type"}}
	}
	return true, nil
}

func validate(id int) (bool, error) {
	if id != 0 {
		return false, errors.InvalidParam{Param: []string{"id"}}
	}
	return true, nil
}
