package products

import (
	"context"
	"fmt"
	"testing"
	"zopsmart/gofr-curd/model"

	//	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/stretchr/testify/assert"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	testAddProduct(t, app)
	testGetProductByID(t, app)
	testAddProductWithError(t, app)
	testUpdateProduct(t, app)
	testGetProducts(t, app)
	testDeleteProduct(t, app)
	testErrors(t, app)
}


var id int

func testAddProduct(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc    string
		product model.Product
		err     error
	}{
		{"create succuss test #1", model.Product{Name: "Test123", Type: "Test"}, nil},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()
		resp, err := store.AddProduct(ctx, tc.product)
		id = resp
		fmt.Println(id)
		app.Logger.Log(resp)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func testAddProductWithError(t *testing.T, app *gofr.Gofr) {
	customer := model.Product{
		Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj",
	}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	_, err := store.AddProduct(ctx, customer)
	if _, ok := err.(errors.DB); err != nil && ok == false {
		t.Errorf("Error Testcase FAILED")
	}
}

func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		id   int
		err  error
	}{
		{"Get existent id", id, nil},
		{"Get non existent id", 1223, errors.EntityNotFound{Entity: "product", ID: "1223"}},
	}
    fmt.Println(id)
	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		_, err := store.GetProductById(ctx, tc.id)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func testUpdateProduct(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc     string
		customer model.Product
		err      error
	}{
		{"update succuss", model.Product{Id: id, Name: "Test1234"}, nil},
		{"update fail", model.Product{Id: 1, Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj"}, errors.DB{}},
	}

	for i, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		_, err := store.UpdateById(ctx, tc.customer)
		if _, ok := err.(errors.DB); err != nil && ok == false {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.err, err, tc.desc)
		}
	}
}

func testGetProducts(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	_, err := store.GetProducts(ctx)
	if err != nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()

	err := store.DeleteById(ctx, id)
	if err != nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}

func testErrors(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	_ = ctx.DB().Close() // close the connection to generate errors

	store := New()

	err := store.DeleteById(ctx, 64)
	if err == nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}

	_, err = store.GetProducts(ctx)
	if err == nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}

func TestNew(t *testing.T) {
	_ = New()
}
