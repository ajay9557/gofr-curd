package products

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

func validateId(id string) (bool, error) {
	if id == "" {
		return false, errors.MissingParam{Param: []string{"id"}}
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		return false, errors.InvalidParam{Param: []string{"id"}}
	}
	return true, nil
}
