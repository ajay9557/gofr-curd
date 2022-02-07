package products

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	// Seeder will populate actual database with default data defined in .csv file.
	seeder := datastore.NewSeeder(&app.DataStore, "../../db")
	// seeder.ResetCounter = true

	seeder.RefreshTables(t, "products")

	db, mock, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	testGetProductByID(t, app, mock)
	testGetProducts(t, app, mock)
	testCreate(t, app, mock)
	testUpdateById(t, app, mock)
	testDeleteById(t, app, mock)
	seeder.RefreshTables(t, "products")

}

func testGetProductByID(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	// rows := mock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")

	tests := []struct {
		desc            string
		id              int
		expectedProduct *models.Product
		err             error
	}{
		{
			desc: "Get existent id",
			id:   1,
			expectedProduct: &models.Product{
				Id:       1,
				Name:     "mouse",
				Category: "electronics",
			},
			err: nil,
		},
		{
			desc:            "Get non existent id",
			id:              100,
			expectedProduct: nil,
			err:             errors.EntityNotFound{Entity: "products", ID: "100"},
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
	tests := []struct {
		desc            string
		expectedProduct []*models.Product
		err             error
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
			err: nil,
		},
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

func testCreate(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	tests := []struct {
		desc          string
		expectedError error
	}{
		{
			desc:          "Success Case",
			expectedError: nil,
		},
	}

	store := New()

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.Create(ctx, models.Product{Id: 2, Name: "volleyball", Category: "sports"})

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testUpdateById(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	tests := []struct {
		desc          string
		id            int
		input         models.Product
		expectedError error
	}{
		{
			desc: "Success Case",
			id:   1,
			input: models.Product{
				Id:       1,
				Name:     "mouse",
				Category: "gaming",
			},
			expectedError: nil,
		},
	}

	store := New()

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.UpdateById(ctx, tc.id, tc.input)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testDeleteById(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	tests := []struct {
		desc          string
		id            int
		expectedError error
	}{
		{
			desc:          "Success Case",
			id:            1,
			expectedError: nil,
		},
	}

	store := New()

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.DeleteById(ctx, tc.id)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}
