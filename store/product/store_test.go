package product

import (
	"context"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/stretchr/testify/assert"
	"github.com/tejas/gofr-crud/models"
)


func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	seeder := datastore.NewSeeder(&app.DataStore, "../db")

	seeder.ResetCounter = true

	testGetProductById(t, app)
	testGetAllProducts(t, app)
	testUpdateProductById(t, app)
	testDeleteProduct(t, app)
	testCreateProductWithoutErr(t, app)
	testCreateProductWithErr(t, app)
}

func testGetProductById(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		id   int
		err  error
	}{
		{
			desc: "Case 1 : Success Case ( existent id )",
			id:   1,
			err:  nil,
		},
		{
			desc: "Case 2 : Failure case ( non existent id )",
			id:   1221,
			err:  errors.EntityNotFound{Entity: "product", ID: "1221"},
		},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)

		ctx.Context = context.Background()

		store := New()

		_, err := store.GetProductById(ctx, tc.id)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func testGetAllProducts(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	_, err := store.GetAllProducts(ctx)

	if err != nil {
		t.Errorf("failed, expected: %v, Got: %v", nil, err)
	}

}

func testUpdateProductById(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc     string
		expected models.Product
		err      error
	}{
		{
			desc:     "Case 1: Success Case",
			expected: models.Product{Id: 1, Name: "product1", Type: "type1"},
			err:      nil,
		},
		{
			desc:     "Case 2: Failure Case",
			expected: models.Product{Id: 1, Name: "very-long-name-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Type: "typee"},
			err:      errors.DB{},
		},
	}

	for i, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		_, err := store.UpdateProduct(ctx, tc.expected)

		if _, ok := err.(errors.DB); err != nil && ok == false {
			t.Errorf("TEST[%v] Failed.\tExpected: %v\t, Got: %v\n", i, tc.err, tc.desc)
		}
	}
}

func testCreateProductWithoutErr(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc    string
		product models.Product
		err     error
	}{
		{
			desc:    "Case 1: Success case 1",
			product: models.Product{Name: "name1", Type: "type1"},
			err:     nil,
		},
		{
			desc:    "Case 2: Success case 2",
			product: models.Product{Name: "name2", Type: "type2"},
			err:     nil,
		},
	}

	for i, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		resp, err := store.CreateProduct(ctx, tc.product)

		app.Logger.Log(resp)

		assert.Equal(t, tc.err, err, "TEST[%d], failed: %v", i, tc.desc )
	}

}

func testCreateProductWithErr(t *testing.T, app *gofr.Gofr) {
	customer := models.Product{
		Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj",
		Type: "typeaa",
	}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	_, err := store.CreateProduct(ctx, customer)
	if _, ok := err.(errors.DB); err != nil && ok == false {
		t.Errorf("Error Testcase FAILED")
	}
}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	err := store.DeleteProduct(ctx, 2)

	if err != nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}

/*
func TestGetProductById(t *testing.T) {

	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal("error in stub database connection")
	}

	defer db.Close()

	database, err := gorm.Open("mysql", db)

	if err != nil {
		fmt.Println("error in gorm database connection")
	}

	app.ORM = database

	testcases := []struct {
		desc        string
		input       int
		expectedErr error
		mockCall    []interface{}
		expected    models.Product
	}{
		{
			desc:        "Case 1: Success Case",
			input:       1,
			expectedErr: nil,
			mockCall: []interface{}{
				mock.ExpectQuery("select id, name, type from product where id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
					AddRow(1, "name-1", "test-1")),
			},
			expected: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
		},
		{
			desc:        "Case 2: Failure Case",
			input:       0,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "0"},
			mockCall: []interface{}{
				mock.ExpectQuery("select id, name, type from product").
					WithArgs(0).
					WillReturnError(errors.EntityNotFound{Entity: "product", ID: "0"}),
			},
			expected: models.Product{},
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.GetProductById(ctx, test.input)
			if err != nil && !reflect.DeepEqual(test.expectedErr, err) {
				fmt.Print("expected ", test.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(test.expected, res) {
				fmt.Print("expected ", test.expected, "obtained", res)
			}
		})
	}

}
*/
