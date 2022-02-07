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
	testInsertProduct(t, app)
	testUpdateProduct(t, app)
	testProductsGetById(t, app)
	testProductDeleteById(t, app)
	testAllProducts(t, app)
}

func testInsertProduct(t *testing.T, app *gofr.Gofr) {

	tcs := []struct {
		desc           string
		input          model.ProductDetails
		err            error
		expectedOutput model.ProductDetails
	}{
		{
			desc: "Success",
			err:  nil,
			input: model.ProductDetails{
				Id:    6,
				Name:  "Philips",
				Types: "Light",
			},
			expectedOutput: model.ProductDetails{
				Id:    6,
				Name:  "Philips",
				Types: "Light",
			},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.CreateProducts(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}

func testUpdateProduct(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc  string
		id    int
		err   error
		input model.ProductDetails
	}{
		{
			desc: "Success",
			id:   6,
			err:  nil,
			input: model.ProductDetails{
				Id:    6,
				Name:  "Philips",
				Types: "Light",
			},
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := store.UpdateProductById(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
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
func testProductDeleteById(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc string
		Id   int
		err  error
	}{
		{
			desc: "Success",
			Id:   6,
			err:  nil,
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			err := store.DeleteProductId(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testAllProducts(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()
	_, err := store.GetAll(ctx)
	if err != nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}
