package products

import (
	"reflect"
	"testing"
	"zopsmart/gofr-curd/model"

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
		{"Case 3", "", false, errors.MissingParam{Param: []string{"id"}}},
		{"Case 4", "-1", false, errors.InvalidParam{Param: []string{"id"}}},
		{"Case 5", "sa", false, errors.InvalidParam{Param: []string{"id"}}},
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

func Test_CheckBody(t *testing.T) {
	Tests := []struct {
		desc   string
		input  model.Product
		output bool
		err    error
	}{
		{"Case 1", model.Product{Id: 0, Name: "sarah", Type: "some"}, true, nil},
		{"Case 2", model.Product{Id: 1, Name: "sarah", Type: "some"}, false, errors.InvalidParam{Param: []string{"id"}}},
		{"Case 3", model.Product{Id: 0, Name: "", Type: "some"}, false, errors.MissingParam{Param: []string{"Name"}}},
		{"Case 4", model.Product{Id: 0, Name: "sarah", Type: ""}, false, errors.MissingParam{Param: []string{"Type"}}},
	}
	for _, tes := range Tests {
		t.Run(tes.desc, func(t *testing.T) {
			output, err := CheckBody(tes.input)
			if output != tes.output {
				t.Errorf("expected %t got %t", tes.output, output)
			}
			if !reflect.DeepEqual(tes.err, err) {
				t.Errorf("expected %s got %s", tes.err, err)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	Tests := []struct {
		desc   string
		input  int
		output bool
		err    error
	}{
		{"Case 1", 0, true, nil},
		{"Case 2", 5, false, errors.InvalidParam{Param: []string{"id"}}},
	}
	for _, tes := range Tests {
		t.Run(tes.desc, func(t *testing.T) {
			output, err := validate(tes.input)
			if output != tes.output {
				t.Errorf("expected %t got %t", tes.output, output)
			}
			if !reflect.DeepEqual(tes.err, err) {
				t.Errorf("expected %s got %s", tes.err, err)
			}
		})
	}
}
