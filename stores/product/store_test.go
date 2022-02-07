package product

import (
	"context"
	"gofr-curd/model"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	testProductsGetById(t, app)

}

func testProductsGetById(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc           string
		Id             int
		err            error
		expectedOutput model.ProductDetails
	}{
		{
			desc: "Success",
			Id:   3,
			err:  nil,
			expectedOutput: model.ProductDetails{
				Id:    3,
				Name:  "Sony",
				Types: "TV",
			},
		},
		{
			desc:           "Failure",
			Id:             0,
			err:            errors.EntityNotFound{Entity: "Product", ID: "0"},
			expectedOutput: model.ProductDetails{},
			//Mock:           nil,
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.GetProductById(tc.Id, ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}
