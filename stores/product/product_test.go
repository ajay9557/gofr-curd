package product

import (
	"context"
	"database/sql"
	goError "errors"
	"log"
	"product/models"
	"product/stores"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func setMocker() (*gofr.Context, stores.Store, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
	}
	application := gofr.New()
	ctx := gofr.NewContext(nil, nil, application)
	ctx.Context = context.Background()
	ctx.DB().DB = db
	st := New()
	return ctx, st, mock, db
}

func TestGetProductById(t *testing.T) {
	ctx, store, mock, db := setMocker()
	defer db.Close()
	tests := []struct {
		desc            string
		input           int
		expectedProduct models.Product
		err             error
		mock            *sqlmock.ExpectedQuery
	}{
		{
			desc:  "Test Case 1",
			input: 1,
			expectedProduct: models.Product{
				ID:   1,
				Name: "mouse",
				Type: "electronics",
			},
			err:  nil,
			mock: mock.ExpectQuery("select * from product where id=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).AddRow(1, "mouse", "electronics")),
		},
	}

	for _, tcs := range tests {
		result, err := store.GetProductById(ctx, tcs.input)
		if !reflect.DeepEqual(err, tcs.err) {
			t.Errorf("%s : expected %v, but got %v", tcs.desc, tcs.err, err)
		}

		if tcs.err == nil && !reflect.DeepEqual(result, tcs.expectedProduct) {
			t.Errorf("%s : expected %v, but got %v", tcs.desc, tcs.err, result)
		}
	}
}

func TestGetAllProduct(t *testing.T) {
	ctx, store, mock, db := setMocker()
	defer db.Close()
	testCases := []struct {
		desc           string
		expectedError  error
		expectedOutput []models.Product
		mock           *sqlmock.ExpectedQuery
	}{
		{
			desc:          "Test Case 1",
			expectedError: nil,
			expectedOutput: []models.Product{{
				ID:   1,
				Name: "item1",
				Type: "type1",
			},
				{
					ID:   2,
					Name: "item2",
					Type: "type2",
				},
			},
			mock: mock.ExpectQuery("select * from product").
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
					AddRow(1, "item1", "type1").
					AddRow(2, "item2", "type2")),
		},
	}

	for _, tcs := range testCases {
		tc := tcs
		out, err := store.GetAllProduct(ctx)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("%s : expected %v, your output %v", tc.desc, tc.expectedError, err)
		}

		if tc.expectedError == nil && !reflect.DeepEqual(out, tc.expectedOutput) {
			t.Errorf("%s : expected %v, your output %v", tc.desc, tc.expectedOutput, out)
		}
	}
}

func TestAddProduct(t *testing.T) {
	ctx, store, mock, db := setMocker()
	defer db.Close()
	testCases := []struct {
		desc          string
		input         models.Product
		expectedError error
		mock          *sqlmock.ExpectedExec
	}{
		{
			desc:          "Test Case 1",
			input:         models.Product{ID: 1, Name: "novo", Type: "trim"},
			expectedError: nil,
			mock: mock.ExpectExec("insert into product(name, type) values(?, ?)").
				WithArgs("novo", "trim").WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Test Case 2",
			input:         models.Product{ID: 1, Name: "", Type: ""},
			expectedError: goError.New("FAILED TO ADD PRODUCT"),
			mock:          nil,
		},
	}

	for _, tcs := range testCases {
		tc := tcs

		err := store.AddProduct(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expectedError, err)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	ctx, store, mock, db := setMocker()
	defer db.Close()
	testCases := []struct {
		desc          string
		input         models.Product
		expectedError error
		mock          *sqlmock.ExpectedExec
	}{
		{
			desc:          "Test Case 1",
			input:         models.Product{ID: 3, Name: "novo", Type: "trimmer"},
			expectedError: nil,
			mock: mock.ExpectExec("update product set name=?, type=? where id=?").
				WithArgs("novo", "trimmer", 3).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Test Case 2",
			input:         models.Product{},
			expectedError: goError.New("FAILED TO UPDATE THE PRODUCT"),
			mock:          nil,
		},
	}

	for _, tcs := range testCases {
		tc := tcs
		err := store.UpdateProduct(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expectedError, err)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	ctx, store, mock, db := setMocker()
	defer db.Close()
	testCases := []struct {
		desc          string
		input         int
		expectedError error
		mock          *sqlmock.ExpectedExec
	}{
		{
			desc:          "Test Case 1",
			input:         2,
			expectedError: nil,
			mock: mock.ExpectExec("delete from product where id=?").
				WithArgs(2).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Test Case 2",
			input:         -1,
			expectedError: goError.New("FAILED TO DELETE PRODUCT"),
			mock:          nil,
		},
	}

	for _, tcs := range testCases {
		err := store.DeleteProduct(ctx, tcs.input)

		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("%s : expected %v, but got %v", tcs.desc, tcs.expectedError, err)
		}
	}
}
