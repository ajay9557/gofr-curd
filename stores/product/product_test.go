package product

import (
	"context"
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	goError "errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"log"
	"product/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestStoreLayer(t *testing.T) {
	application := gofr.New()
	seeder := datastore.NewSeeder(&application.DataStore, "../../db")
	seeder.RefreshTables(t, "product")

	db, _, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	application.ORM = database
	testGetProductById(t, application)
	testGetAllProduct(t, application)
	testAddProduct(t, application)
	testUpdateProduct(t, application)
	testDeleteProduct(t, application)
}

func testGetProductById(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc            string
		id              int
		expectedProduct models.Product
		err             error
	}{
		{
			desc: "Get existent id",
			id:   1,
			expectedProduct: models.Product{
				Id:   1,
				Name: "mouse",
				Type: "electronics",
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			_, err := store.GetProductById(ctx, tc.id)
			if err != nil {
				t.Errorf("Expected: %s, Ouput: %s", tc.err, err)
			}
		})
	}
}

func testGetAllProduct(t *testing.T, application *gofr.Gofr) {
	testCases := []struct {
		desc          string
		expectedError error
	}{
		{
			desc:          "Test Case 1",
			expectedError: nil,
		},
	}

	for _, tcs := range testCases {
		ctx := gofr.NewContext(nil, nil, application)
		ctx.Context = context.Background()
		store := New()
		_, err := store.GetAllProduct(ctx)
		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Expected: %s, Result: %s", tcs.expectedError, err)
		}
	}
}

func testAddProduct(t *testing.T, application *gofr.Gofr) {
	testCases := []struct {
		desc          string
		input         models.Product
		expectedError error
	}{
		{
			desc:          "Test Case 1",
			input:         models.Product{Id: 1, Name: "novo", Type: "Trimmer"},
			expectedError: nil,
		},
		{
			desc:          "Test Case 2",
			input:         models.Product{Id: 1, Name: "", Type: ""},
			expectedError: goError.New("FAILED TO ADD PRODUCT"),
		},
	}

	for _, tcs := range testCases {
		ctx := gofr.NewContext(nil, nil, application)
		ctx.Context = context.Background()
		store := New()
		err := store.AddProduct(ctx, tcs.input)
		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Expected: %s, Result: %s", tcs.expectedError, err)
		}
	}
}

func testUpdateProduct(t *testing.T, application *gofr.Gofr) {
	testCases := []struct {
		desc          string
		input         models.Product
		expectedError error
	}{
		{
			desc:          "Test Case 1",
			input:         models.Product{Id: 9, Name: "novo", Type: "trimmer"},
			expectedError: nil,
		},
		{
			desc:          "Test Case 2",
			input:         models.Product{},
			expectedError: goError.New("FAILED TO UPDATE THE PRODUCT"),
		},
	}

	for _, tcs := range testCases {
		ctx := gofr.NewContext(nil, nil, application)
		ctx.Context = context.Background()
		store := New()
		err := store.UpdateProduct(ctx, tcs.input)
		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Expected: %s, Result: %s", tcs.expectedError, err)
		}
	}
}

func testDeleteProduct(t *testing.T, application *gofr.Gofr) {
	testCases := []struct {
		desc          string
		input         int
		expectedError error
	}{
		{
			desc:          "Test Case 1",
			input:         1,
			expectedError: nil,
		},
		{
			desc:          "Test Case 2",
			input:         -1,
			expectedError: goError.New("FAILED TO DELETE PRODUCT"),
		},
	}

	for _, tcs := range testCases {
		ctx := gofr.NewContext(nil, nil, application)
		ctx.Context = context.Background()
		store := New()
		err := store.DeleteProduct(ctx, tcs.input)
		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Expected: %s, Result: %s", tcs.expectedError, err)
		}
	}
}
