package product

import (
	"context"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
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
	db, mock, _ := sqlmock.New()

	defer db.Close()

	tests := []struct {
		desc     string
		ID       int
		expected models.Product
		err      error
		mockCall []interface{}
	}{
		{
			desc:     "Case 1 : Success Case",
			ID:       1,
			expected: models.Product{ID: 1, Name: "product1", Type: "type1"},
			err:      nil,
			mockCall: []interface{}{mock.ExpectQuery(`select id, name, type from product where id = ?`).
				WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "Type"}).AddRow(1, "product1", "type1"))},
		},
		{
			desc:     "Case 2 : Failure case",
			ID:       1221,
			expected: models.Product{},
			err:      errors.Error("product data not found for the given id"),
			mockCall: []interface{}{mock.ExpectExec(`select id, name, type from product where id = ?`).
				WillReturnError(errors.Error("product data not found for the given id"))},
		},
	}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	for _, test := range tests {
		ts := test

		t.Run(ts.desc, func(t *testing.T) {
			resp, err := store.GetProductByID(ctx, ts.ID)

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
	db, mock, _ := sqlmock.New()

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)

	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	testCases := []struct {
		desc     string
		expOut   []models.Product
		expErr   error
		mockCall []interface{}
	}{
		{
			desc: "Case 1: Success Case",
			expOut: []models.Product{
				{
					ID:   1,
					Name: "name-1",
					Type: "type-1",
				},
			},
			expErr: nil,
			mockCall: []interface{}{mock.ExpectQuery(`select id, name, type from product`).
				WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "Type"}).AddRow(1, "name-1", "type-1"))},
		},
		{
			desc:   "Case 1: Failure Case",
			expOut: nil,
			expErr: errors.Error("internal db error"),
			mockCall: []interface{}{mock.ExpectQuery(`select id, name, type from product`).
				WillReturnError(errors.Error("internal db error"))},
		},
	}

	for _, test := range testCases {
		ts := test

		t.Run(ts.desc, func(t *testing.T) {
			res, err := store.GetAllProducts(ctx)

			if !reflect.DeepEqual(err, ts.expErr) {
				t.Errorf("Expected :%v, Got : %v", ts.expErr, err)
			}

			if !reflect.DeepEqual(res, ts.expOut) {
				t.Errorf("Expected :%v, Got : %v", ts.expOut, res)
			}
		})
	}
}

func testUpdateProductByID(t *testing.T, app *gofr.Gofr) {
	db, mock, _ := sqlmock.New()

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	testCases := []struct {
		desc     string
		ID       int
		expected models.Product
		err      error
		mockCall []interface{}
	}{
		{
			desc:     "Case 1: Success Case",
			ID:       1,
			expected: models.Product{ID: 1, Name: "product1", Type: "type1"},
			err:      nil,
			mockCall: []interface{}{mock.ExpectExec(`update product set`).
				WithArgs("product1", "type1", 1).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
		{
			desc: "Case 2: Failure Case",
			expected: models.Product{
				ID:   2,
				Name: "very-long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Type: "typee"},
			err: errors.Error("error while updating product"),
			mockCall: []interface{}{mock.ExpectExec(`update product set`).
				WithArgs("name-2", "type-2", 2).WillReturnError(errors.Error("error while updating product"))},
		},
	}

	for _, test := range testCases {
		ts := test

		t.Run(ts.desc, func(t *testing.T) {
			_, err := store.UpdateProduct(ctx, ts.expected)
			if !reflect.DeepEqual(ts.err, err) {
				t.Errorf("Expected: %v, Got: %v", ts.err, err)
			}
		})
	}
}

func testCreateProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, _ := sqlmock.New()

	defer db.Close()
	
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	testCases := []struct {
		desc     string
		input    models.Product
		err      error
		mockCall []interface{}
	}{
		{
			desc:  "Case 1: Success case 1",
			input: models.Product{ID: 1, Name: "name-1", Type: "type-1"},
			err:   nil,
			mockCall: []interface{}{mock.ExpectExec(`insert into product values`).
				WithArgs(1, "name-1", "type-1").WillReturnResult(sqlmock.NewResult(1, 1))},
		},
		{
			desc: "Case 2: Failure Case",
			input: models.Product{
				ID:   2,
				Name: "long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Type: "typeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			},
			err: errors.Error("internal db error"),
			mockCall: []interface{}{mock.ExpectExec(`insert into product values`).
				WillReturnError(errors.Error("internal db error"))},
		},
	}

	for _, test := range testCases {
		ts := test

		t.Run(ts.desc, func(t *testing.T) {
			_, err := store.CreateProduct(ctx, ts.input)

			if !reflect.DeepEqual(err, ts.err) {
				t.Errorf("Expected :%v, Got : %v", ts.err, err)
			}
		})
	}
}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, _ := sqlmock.New()

	defer db.Close()
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	testCases := []struct {
		desc     string
		ID       int
		expErr   error
		mockCall []interface{}
	}{
		{
			desc:   "Case 1: Success Case",
			ID:     1,
			expErr: nil,
			mockCall: []interface{}{mock.ExpectExec(`delete from product where id = ?`).
				WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
		{
			desc:   "Case 2: Failure Case",
			ID:     1,
			expErr: errors.Error("internal db error"),
			mockCall: []interface{}{mock.ExpectExec(`delete from product where id = ?`).
				WillReturnError(errors.Error("internal db error"))},
		},
	}

	for _, test := range testCases {
		ts := test

		t.Run(ts.desc, func(t *testing.T) {
			err := store.DeleteProduct(ctx, ts.ID)

			if !reflect.DeepEqual(err, ts.expErr) {
				t.Errorf("Expected :%v, Got : %v", ts.expErr, err)
			}
		})
	}
}
