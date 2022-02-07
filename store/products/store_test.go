package products

import (
	"context"
	"testing"

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
	seeder.RefreshTables(t, "products")

	testGetProductByID(t, app)
	testGetProducts(t, app)
	testCreate(t, app)
	testUpdateByID(t, app)
	testDeleteByID(t, app)
	seeder.RefreshTables(t, "products")
}

func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc            string
		id              int
		expectedProduct *models.Product
		expectedError   error
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
		},
		{
			desc:            "Get non existent id",
			id:              100,
			expectedProduct: nil,
			expectedError:   errors.EntityNotFound{Entity: "products", ID: "100"},
		},
	}

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			p, err := store.GetByID(ctx, tc.id)
			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedProduct, p, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testGetProducts(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc            string
		expectedProduct []*models.Product
		err             error
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
			err: nil,
		},
	}

	for index, test := range tests {
		i := index
		tc := test

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

func testCreate(t *testing.T, app *gofr.Gofr) {
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

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.Create(ctx, models.Product{ID: 2, Name: "volleyball", Category: "sports"})

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testUpdateByID(t *testing.T, app *gofr.Gofr) {
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
				ID:       1,
				Name:     "mouse",
				Category: "gaming",
			},
			expectedError: nil,
		},
	}

	store := New()

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.UpdateByID(ctx, tc.id, tc.input)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

func testDeleteByID(t *testing.T, app *gofr.Gofr) {
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

	for index, test := range tests {
		i := index
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			err := store.DeleteByID(ctx, tc.id)

			assert.Equal(t, tc.expectedError, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}
