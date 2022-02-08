package products

import (
	"context"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	testInsertProduct(t, app)
	testUpdateProduct(t, app)
	testProductsGetByID(t, app)
	testProductDeleteByID(t, app)
	testAllProducts(t, app)
}
func testInsertProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcs := []struct {
		desc  string
		input models.Product
		err   error
		Mock  []interface{}
	}{
		{
			desc: "Success",
			err:  nil,
			input: models.Product{
				ID:   4,
				Name: "Brand",
				Type: "Mafti",
			},
			Mock: []interface{}{
				mock.ExpectExec(`Insert into Product values`).WithArgs(4, "Brand", "Mafti").
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc: "Failure",
			err:  errors.Error("Internal DB error"),
			input: models.Product{
				ID:   3,
				Name: "very-long-name",
				Type: "dummy",
			},
			Mock: []interface{}{
				mock.ExpectExec(`Insert into Product values`).WillReturnError(errors.Error("Internal DB error")),
			},
		},
	}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			_, err := store.CreateProducts(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testUpdateProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcs := []struct {
		desc  string
		ID    int
		err   error
		input models.Product
		Mock  []interface{}
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
			input: models.Product{
				ID:   6,
				Name: "Brands",
				Type: "Twills",
			},
			Mock: []interface{}{
				mock.ExpectExec(`Update Product set`).WithArgs("Brands", "Twills", 6).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc: "Failure case",
			ID:   3,
			input: models.Product{
				ID:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoia",
				Type: "food",
			},
			err: errors.Error("Internal DB error"),
			Mock: []interface{}{
				mock.ExpectExec(`Update Product set`).WithArgs("Brands", "Twills", 6).WillReturnError(errors.Error("Internal DB error")),
			},
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			err := store.UpdateID(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testProductsGetByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	defer db.Close()

	query := `Select ID,Name,Type from Product where Id =?`

	tcs := []struct {
		desc           string
		ID             int
		err            error
		expectedOutput models.Product
		Mock           []interface{}
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
			expectedOutput: models.Product{
				ID:   6,
				Name: "Brands",
				Type: "Twills",
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows(
				[]string{"ID", "Name", "Type"}).AddRow(6, "Brands", "Twills"))},
		},
		{
			desc:           "Failure",
			ID:             0,
			err:            errors.EntityNotFound{Entity: "Product", ID: "0"},
			expectedOutput: models.Product{},
			Mock: []interface{}{
				mock.ExpectExec(query).WillReturnError(errors.EntityNotFound{Entity: "Product", ID: "0"}),
			},
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	ctx.DB().DB = db

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.GetID(ctx, tc.ID)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}
func testProductDeleteByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	defer db.Close()

	tcs := []struct {
		desc string
		ID   int
		err  error
		Mock []interface{}
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
			Mock: []interface{}{
				mock.ExpectExec(`Delete from Product where id=?`).WithArgs(6).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}

	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		ctx.DB().DB = db
		t.Run(tc.desc, func(t *testing.T) {
			err := store.DeleteID(ctx, tc.ID)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testAllProducts(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	defer db.Close()

	query := `Select Id,Name,Type from Product`
	tcs := []struct {
		desc           string
		expectedOutput []models.Product
		err            error
		Mock           []interface{}
	}{
		{
			desc: "Success",
			expectedOutput: []models.Product{
				{
					ID:   1,
					Name: "Shirts",
					Type: "Jeans",
				},
			},
			err: nil,
			Mock: []interface{}{
				mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows(
					[]string{"ID", "Name", "Type"}).AddRow(1, "Shirts", "Jeans")),
			},
		},
		{
			desc:           "Failure",
			expectedOutput: nil,
			err:            errors.Error("Internal DB error"),
			Mock: []interface{}{
				mock.ExpectQuery(query).WillReturnError(errors.Error("Internal DB error")),
			},
		},
		// {
		// 	desc:           "Failure-1",
		// 	expectedOutput: nil,
		// 	err:           errors.Error("Error in scanning the attributes"),
		// 	Mock: []interface{}{
		// 		mock.ExpectQuery(query).WillReturnError(errors.Error("Error in scanning the attributes")),
		// 	},
		// },
	}

	//
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		ctx.DB().DB = db

		t.Run(tc.desc, func(*testing.T) {
			res, err := store.GetAll(ctx)

			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}

			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}
