package product

import (
	"context"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testCases := []struct {
		desc        string
		id          int
		expectedErr error
		Mock        []interface{}
		expectedRes models.Product
	}{
		{
			desc:        "Success case",
			id:          1,
			expectedErr: nil,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "name-1", "type-1")),
			},
			expectedRes: models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			},
		},
		{
			desc: "Failure case",
			id:   1234,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1234).
					WillReturnError(errors.EntityNotFound{Entity: "product", ID: "1234"}),
			},
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "1234"},
			expectedRes: models.Product{},
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db
			store := New()
			res, err := store.GetByID(ts.id, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				t.Error("expected ", ts.expectedRes, "obtained", res)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testCases := []struct {
		desc        string
		expectedErr error
		Mock        []interface{}
		expectedRes []models.Product
	}{
		{
			desc: "Success case",
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "name-2", "type-2")),
			},
			expectedErr: nil,
			expectedRes: []models.Product{
				{
					ID:   1,
					Name: "name-2",
					Type: "type-2",
				},
			},
		},
		{
			desc:        "Failure case",
			expectedErr: errors.Error("internal db error"),
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").
					WillReturnError(errors.Error("internal db error")),
			},
			expectedRes: nil,
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db
			store := New()
			res, err := store.GetAllProducts(ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				t.Error("expected ", ts.expectedRes, "obtained", res)
			}
		})
	}
}

func TestInsertProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testCases := []struct {
		desc        string
		Mock        []interface{}
		expectedErr error
		product     models.Product
	}{
		{
			desc: "Success case",
			Mock: []interface{}{
				mock.ExpectExec("insert into product").WithArgs(7, "biryani", "food").
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expectedErr: nil,
			product: models.Product{
				ID:   7,
				Name: "biryani",
				Type: "food",
			},
		},
		{
			desc:        "Failure case",
			expectedErr: errors.Error("Error in executing query"),
			Mock: []interface{}{
				mock.ExpectExec("insert into product").WithArgs(2, "very-long-mock-name", "food").
					WillReturnError(errors.Error("Error in executing query")),
			},
			product: models.Product{
				ID:   2,
				Name: "very-long-mock-name",
				Type: "food",
			},
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db
			store := New()
			err := store.InsertProduct(ts.product, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testCases := []struct {
		desc        string
		product     models.Product
		Mock        []interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			product: models.Product{
				ID:   7,
				Name: "apple",
				Type: "food",
			},
			Mock: []interface{}{
				mock.ExpectExec("update product set").WithArgs("apple", "food", 7).
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			product: models.Product{
				ID:   2,
				Name: "very-long-mock-name",
				Type: "food",
			},
			Mock: []interface{}{
				mock.ExpectExec("update product set").WithArgs("very-long-mock-name", "food", 2).
					WillReturnError(errors.Error("Error in executing query")),
			},
			expectedErr: errors.Error("Error in executing query"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db
			store := New()
			err := store.UpdateProduct(ts.product, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestDeleteById(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testCases := []struct {
		desc        string
		id          int
		Mock        []interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   1,
			Mock: []interface{}{
				mock.ExpectExec("delete from product where Id=").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			id:   7,
			Mock: []interface{}{
				mock.ExpectExec("delete from product where Id=").WithArgs(7).
					WillReturnError(errors.Error("internal db error")),
			},
			expectedErr: errors.Error("internal db error"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db
			store := New()
			err := store.DeleteByID(ts.id, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}
