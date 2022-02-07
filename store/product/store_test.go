package product

import (
	"context"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestGetById(t *testing.T) {
	app := gofr.New()
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	testCases := []struct {
		desc        string
		id          int
		expectedErr error
		expectedRes models.Product
	}{
		{
			desc:        "Success case",
			id:          1,
			expectedErr: nil,
			expectedRes: models.Product{
				ID:   1,
				Name: "jeans",
				Type: "clothes",
			},
		},
		{
			desc:        "Failure case",
			id:          1234,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "1234"},
			expectedRes: models.Product{},
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
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
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	_, err := store.GetAllProducts(ctx)

	if err != nil {
		t.Errorf("Failed, Expected %v Obtained %v", nil, err)
	}
}

func TestInsertProduct(t *testing.T) {
	app := gofr.New()
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	testCases := []struct {
		desc        string
		expectedErr error
		product     models.Product
	}{
		{
			desc:        "Success case",
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
			product: models.Product{
				ID:   2,
				Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj",
				Type: "food",
			},
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
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
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true
	testCases := []struct {
		desc        string
		product     models.Product
		expectedErr error
	}{
		{
			desc: "Success case",
			product: models.Product{
				ID:   7,
				Name: "apple",
				Type: "food",
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			product: models.Product{
				ID:   2,
				Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj",
				Type: "food",
			},
			expectedErr: errors.Error("Error in executing query"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
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
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true
	testCases := []struct {
		desc        string
		id          int
		expectedErr error
	}{
		{
			desc:        "Success case",
			id:          7,
			expectedErr: nil,
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			err := store.DeleteByID(ts.id, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}
