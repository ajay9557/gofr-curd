package product

import (
	"context"
	"reflect"

	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	testGetProductByID(t, app)
	testGetAllProducts(t, app)
	testUpdateProductByID(t, app)
	testDeleteProduct(t, app)
	testCreateProduct(t, app)
}

func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc     string
		id       int
		expected models.Product
		err      error
	}{
		{
			desc:     "Case 1 : Success Case ( existent id )",
			id:       1,
			expected: models.Product{ID: 1, Name: "product1", Type: "type1"},
			err:      nil,
		},
		{
			desc:     "Case 2 : Failure case ( non existent id )",
			id:       1221,
			expected: models.Product{},
			err:      errors.EntityNotFound{Entity: "product", ID: "1221"},
		},
	}

	for _, test := range tests {
		ts := test
		ctx := gofr.NewContext(nil, nil, app)

		ctx.Context = context.Background()

		store := New()

		t.Run(ts.desc, func(t *testing.T) {
			resp, err := store.GetProductByID(ctx, ts.id)

			if !reflect.DeepEqual(err, ts.err) {
				t.Errorf("Expected :%v, Got : %v", ts.err, err)
			}

			if !reflect.DeepEqual(resp, ts.expected) {
				t.Errorf("Expected :%v, Got : %v", ts.expected, resp)
			}
		})
	}
}

func testGetAllProducts(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	_, err := store.GetAllProducts(ctx)

	if err != nil {
		t.Errorf("expected: %v, Got: %v", nil, err)
	}
}

func testUpdateProductByID(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc     string
		expected models.Product
		err      error
	}{
		{
			desc:     "Case 1: Success Case",
			expected: models.Product{ID: 1, Name: "product1", Type: "type1"},
			err:      nil,
		},
		{
			desc: "Case 2: Failure Case",
			expected: models.Product{ID: 1,
				Name: "very-long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Type: "typee"},
			err: errors.DB{},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		_, err := store.UpdateProduct(ctx, tc.expected)

		if _, ok := err.(errors.DB); err != nil && ok == false {
			t.Errorf("Expected: %v, Got: %v", tc.err, tc.desc)
		}
	}
}

func testCreateProduct(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		input  models.Product
		expOut models.Product
		err    error
	}{
		{
			desc:   "Case 1: Success case 1",
			input:  models.Product{ID: 10, Name: "name1", Type: "type1"},
			expOut: models.Product{ID: 10, Name: "name1", Type: "type1"},
			err:    nil,
		},
		{
			desc: "Case 2: Failure Case",
			input: models.Product{
				ID:   2,
				Name: "long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Type: "typeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			},
			expOut: models.Product{
				ID:   2,
				Name: "long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Type: "typeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			},
			err: errors.Error("error in inserting new product"),
		},
	}

	for _, test := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ts := test

		store := New()

		t.Run(ts.desc, func(t *testing.T) {
			res, err := store.CreateProduct(ctx, ts.input)

			if !reflect.DeepEqual(err, ts.err) {
				t.Errorf("Expected :%v, Got : %v", ts.err, err)
			}
			if !reflect.DeepEqual(res, ts.expOut) {
				t.Errorf("Expected :%v, Got : %v", ts.expOut, res)
			}
		})
	}
}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		id     int
		expOut models.Product
		expErr error
	}{
		{
			desc: "Case 1: Success Case",
			id:   10,
			expOut: models.Product{
				ID: 10, Name: "name1", Type: "type1",
			},
			expErr: nil,
		},
	}

	for _, test := range testCases {
		ts := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		t.Run(ts.desc, func(t *testing.T) {
			err := store.DeleteProduct(ctx, ts.id)

			if !reflect.DeepEqual(err, ts.expErr) {
				t.Errorf("Expected :%v, Got : %v", ts.expErr, err)
			}
		})
	}
}
