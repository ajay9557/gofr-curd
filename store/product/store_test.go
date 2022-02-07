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

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	seeder := datastore.NewSeeder(&app.DataStore, "../../db")
	seeder.ResetCounter = true
	//testGet(t, app)
	seeder.RefreshTables(t, "products")
	//testGetById(t, app)
	testCreate(t, app)
	//seeder.RefreshTables(t, "products")
}

func testGet(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		expErr error
		expOut []*models.Product
	}{
		{
			desc:   "success case",
			expErr: nil,
			expOut: []*models.Product{
				&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				},
				&models.Product{
					Id:   2,
					Name: "this",
					Type: "that",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()
			out, err := store.Get(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
			}
		})
	}
}

func testGetById(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		input  int
		expErr error
		expOut *models.Product
		//mockCalls []*sqlmock.ExpectedQuery
	}{
		{
			desc:   "success case",
			input:  1,
			expErr: nil,
			expOut: &models.Product{
				Id:   1,
				Name: "test",
				Type: "example",
			},
			//mockCalls: []*sqlmock.ExpectedQuery{
			//	mock.ExpectQuery("select name, type from products where id=?").
			//		WithArgs(1).
			//		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
			//			AddRow(1, "ss", "example")),
			//},
		},
		{
			desc:  "entity not in database",
			input: 1022,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1022",
			},
			//mockCalls: []*sqlmock.ExpectedQuery{
			//	mock.ExpectQuery("select name, type from products where id=?").
			//		WithArgs(1022).
			//		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"})),
			//},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.GetById(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testCreate(t *testing.T, app *gofr.Gofr) {
	tesCases := []struct {
		desc   string
		input  models.Product
		expErr error
	}{
		{
			desc: "success case",
			input: models.Product{
				Id:   3,
				Name: "this",
				Type: "that",
			},
			expErr: nil,
		},
		{
			desc: "Entity already exists",
			input: models.Product{
				Id:   1,
				Name: "test",
				Type: "example",
			},
			expErr: errors.EntityAlreadyExists{},
		},
	}

	for _, tc := range tesCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		resp, err := store.Create(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, &tc.input) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, &tc.input, resp)
		}

	}
}
