package products

import (
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

func Test_Validateid(t *testing.T) {
	Tests := []struct {
		desc   string
		input  string
		output bool
		err    error
	}{
		{"Case 1", "5", true, nil},
		{"Case 2", "Nogmail", false, errors.InvalidParam{Param: []string{"id"}}},
		{"Case 1", "", false, errors.MissingParam{Param: []string{"id"}}},
	}
	for _, tes := range Tests {
		t.Run(tes.desc, func(t *testing.T) {
			output, err := validateId(tes.input)
			if output != tes.output {
				t.Errorf("expected %t got %t", tes.output, output)
			}
			if !reflect.DeepEqual(tes.err, err) {
				t.Errorf("expected %s got %s", tes.err, err)
			}
		})
	}
}
