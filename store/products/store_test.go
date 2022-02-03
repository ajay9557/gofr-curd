package products

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	db, mock, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	testGetProductByID(t, app, mock)
	testGetProducts(t, app, mock)
}

func testGetProductByID(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	rows := mock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")

	tests := []struct {
		desc            string
		id              int
		expectedProduct *models.Product
		err             error
		mockCall        *sqlmock.ExpectedQuery
	}{
		{
			desc: "Get existent id",
			id:   1,
			expectedProduct: &models.Product{
				Id:       1,
				Name:     "mouse",
				Category: "electronics",
			},
			err:      nil,
			mockCall: mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows),
		},
		{
			desc:            "Get non existent id",
			id:              100,
			expectedProduct: nil,
			err:             errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall:        mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(100).WillReturnRows(mock.NewRows([]string{"id", "name", "category"})),
		},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			p, err := store.GetById(ctx, tc.id)
			assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedProduct, p, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testGetProducts(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	rows := mock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")

	tests := []struct {
		desc            string
		expectedProduct []*models.Product
		err             error
		mockCall        *sqlmock.ExpectedQuery
	}{
		{
			desc: "Get All Products",
			expectedProduct: []*models.Product{
				{
					Id:       1,
					Name:     "mouse",
					Category: "electronics",
				},
			},
			err:      nil,
			mockCall: mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows),
		},
		// {
		// 	desc:            "Connection error/ Context cancelled",
		// 	expectedProduct: nil,
		// 	err:             errors.EntityNotFound{Entity: "products"},
		// 	isCtxDone:       true,
		// 	mockCall:        mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(100).WillReturnRows(mock.NewRows([]string{"id", "name", "category"})),
		// },
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			p, err := store.Get(ctx)
			assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedProduct, p, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}
