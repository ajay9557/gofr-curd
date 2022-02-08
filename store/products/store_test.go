package products

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func getTestData() (*gofr.Gofr, sqlmock.Sqlmock, *sql.DB) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	return app, mock, db
}

func TestGetProductByID(t *testing.T) {
	app, mock, db := getTestData()

	rows := sqlmock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")

	tests := []struct {
		desc            string
		id              int
		expectedProduct *models.Product
		expectedError   error
		mockQuery       *sqlmock.ExpectedQuery
	}{
		{
			desc: "Get existent id",
			id:   1,
			expectedProduct: &models.Product{
				ID:       1,
				Name:     "mouse",
				Category: "electronics",
			},
			expectedError: nil,
			mockQuery:     mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows),
		},
		{
			desc:            "Get non existent id",
			id:              100,
			expectedProduct: nil,
			expectedError:   sql.ErrNoRows,
			mockQuery: mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(100).WillReturnError(
				sql.ErrNoRows,
			),
		},
	}

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db

			store := New()

			p, err := store.GetByID(ctx, tc.id)
			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedProduct, p, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func TestGetProducts(t *testing.T) {
	app, mock, db := getTestData()

	rows := sqlmock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")
	row2 := sqlmock.NewRows([]string{"name", "category"}).AddRow("mouse", "electronics")

	tests := []struct {
		desc            string
		expectedProduct []*models.Product
		err             error
		mockQuery       *sqlmock.ExpectedQuery
	}{
		{
			desc: "Get All Products",
			expectedProduct: []*models.Product{
				{
					ID:       1,
					Name:     "mouse",
					Category: "electronics",
				},
			},
			err:       nil,
			mockQuery: mock.ExpectQuery("SELECT * FROM products").WillReturnRows(rows),
		},
		{
			desc:            "Error while fetching products",
			expectedProduct: nil,
			err:             errors.DB{Err: errors.Error("Error while fetching products")},
			mockQuery: mock.ExpectQuery("SELECT * FROM products").WillReturnError(
				errors.DB{Err: errors.Error("Error while fetching products")},
			),
		},
		{
			desc:            "No record found",
			expectedProduct: nil,
			err:             errors.EntityNotFound{Entity: "product"},
			mockQuery:       mock.ExpectQuery("SELECT * FROM products").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "category"})),
		},
		{
			desc:            "No record found",
			expectedProduct: nil,
			err:             errors.EntityNotFound{Entity: "product"},
			mockQuery:       mock.ExpectQuery("SELECT * FROM products").WillReturnRows(row2),
		},
	}

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db

			store := New()

			p, err := store.Get(ctx)
			assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedProduct, p, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func TestCreate(t *testing.T) {
	app, mock, db := getTestData()

	pr := models.Product{
		ID:       1,
		Name:     "mouse",
		Category: "electronics",
	}

	tests := []struct {
		desc          string
		input         models.Product
		expectedError error
		mockQuery     *sqlmock.ExpectedExec
	}{
		{
			desc:          "Success Case",
			input:         pr,
			expectedError: nil,
			mockQuery: mock.ExpectExec("INSERT INTO products(id, name, category) values(?, ?, ?)").WithArgs(
				pr.ID, pr.Name, pr.Category,
			).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Error while creating product",
			input:         pr,
			expectedError: errors.Error("Error while creating product"),
			mockQuery: mock.ExpectExec("INSERT INTO products(id, name, category) values(?, ?, ?)").WithArgs(
				pr.ID, pr.Name, pr.Category,
			).WillReturnError(errors.Error("Error while creating product")),
		},
	}

	store := New()

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db

			err := store.Create(ctx, tc.input)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func TestUpdateByID(t *testing.T) {
	app, mock, db := getTestData()

	pr := models.Product{
		Name:     "mouse",
		Category: "electronics",
	}

	tests := []struct {
		desc          string
		id            int
		input         models.Product
		expectedError error
		mockQuery     *sqlmock.ExpectedExec
	}{
		{
			desc:          "Success Case",
			id:            1,
			input:         pr,
			expectedError: nil,
			mockQuery: mock.ExpectExec("UPDATE products SET name = ?, category = ? WHERE id = ?").WithArgs(
				pr.Name, pr.Category, 1,
			).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Error while updating product",
			id:            1,
			input:         pr,
			expectedError: errors.Error("Error while updating product"),
			mockQuery: mock.ExpectExec("UPDATE products SET name = ?, category = ? WHERE id = ?").WithArgs(
				pr.Name, pr.Category, 1,
			).WillReturnError(errors.Error("Error while updating product")),
		},
	}

	store := New()

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db

			err := store.UpdateByID(ctx, tc.id, tc.input)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func TestDeleteByID(t *testing.T) {
	app, mock, db := getTestData()

	tests := []struct {
		desc          string
		id            int
		expectedError error
		mockQuery     *sqlmock.ExpectedExec
	}{
		{
			desc:          "Success Case",
			id:            1,
			expectedError: nil,
			mockQuery:     mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Invalid Id",
			id:            100,
			expectedError: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockQuery:     mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(100).WillReturnResult(sqlmock.NewResult(0, 0)),
		},
		{
			desc:          "Error while deleting product",
			id:            1,
			expectedError: sql.ErrConnDone,
			mockQuery:     mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(1).WillReturnError(sql.ErrConnDone),
		},
	}

	store := New()

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			ctx.DB().DB = db

			err := store.DeleteByID(ctx, tc.id)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}
