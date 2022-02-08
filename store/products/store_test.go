package products

import (
	"context"
	"reflect"
	"testing"
	"zopsmart/gofr-curd/model"

	e "errors"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	testAddProduct(t, app)
	testGetProductByID(t, app)
	testUpdateProduct(t, app)
	testGetProducts(t, app)
	testDeleteProduct(t, app)
}

func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db

	rows := sqlmock.NewRows([]string{"Id", "Name", "Type"}).
		AddRow(1, "Sarah Vaughan", "Sarah")

	tests := []struct {
		desc   string
		id     int
		output model.Product
		err    error
		mock   []interface{}
	}{
		{
			desc:   "case-1",
			id:     1,
			output: model.Product{ID: 1, Name: "Sarah Vaughan", Type: "Sarah"},
			err:    nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from Products where id=?").WithArgs(1).WillReturnRows(rows),
			},
		},
		{
			desc:   "case-2",
			id:     1223,
			output: model.Product{},
			err:    errors.EntityNotFound{Entity: "product", ID: "1223"},
			mock: []interface{}{
				mock.ExpectQuery("Select * from Products where id=?").WithArgs(1223).
					WillReturnError(errors.EntityNotFound{Entity: "product", ID: "1223"}),
			},
		},
	}

	for _, tests := range tests {
		tc := tests
		store := New()

		output, err := store.GetProductByID(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %s, got error: %s", tc.err, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("expected %v, got :  %v", tc.output, output)
		}
	}
}

func testAddProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db

	tests := []struct {
		desc   string
		input  model.Product
		output int
		err    error
		mock   []interface{}
	}{
		{
			desc:   "case-1",
			input:  model.Product{Name: "Sarah Vaughan", Type: "Sarah"},
			output: 1,
			err:    nil,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO Products(Id,Name,Type) VALUES(?,?,?)").WithArgs(0, "Sarah Vaughan", "Sarah").
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:   "case-2",
			input:  model.Product{Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkjthuijnhvbjiommjgfbnu", Type: "er"},
			output: -1,
			err:    e.New("error"),
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO Products(Id,Name,Type) VALUES(?,?,?)").
					WithArgs(0, "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkjthuijnhvbjiommjgfbnu", "er").WillReturnError(e.New("error")),
			},
		},
	}

	for _, tests := range tests {
		tc := tests
		store := New()

		output, err := store.AddProduct(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %v, got error: %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("expected %v, got :  %v", tc.output, output)
		}
	}
}

func testGetProducts(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db

	rows := sqlmock.NewRows([]string{"Id", "Name", "Type"}).
		AddRow(1, "Sarah Vaughan", "Sarah")

	tests := []struct {
		desc   string
		output []model.Product
		err    error
		mock   []interface{}
	}{
		{
			desc: "case-1",

			output: []model.Product{{ID: 1, Name: "Sarah Vaughan", Type: "Sarah"}},
			err:    nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from Products").WillReturnRows(rows),
			},
		},
		{
			desc:   "case-2",
			output: nil,
			err:    errors.Error("error"),
			mock: []interface{}{
				mock.ExpectQuery("Select * from Products").WillReturnError(errors.Error("error")),
			},
		},
	}

	for _, tests := range tests {
		tc := tests
		store := New()

		output, err := store.GetProducts(ctx)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %v, got error: %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("expected %v, got :  %v", tc.output, output)
		}
	}
}

func testUpdateProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db

	tests := []struct {
		desc   string
		output model.Product
		err    error
		mock   []interface{}
	}{
		{
			desc:   "case-1",
			output: model.Product{ID: 1, Name: "Sarah Vaughan", Type: "Sarah"},
			err:    nil,
			mock: []interface{}{
				mock.ExpectExec("Update Products set Name=?,Type=? where Id=?").WithArgs("Sarah Vaughan", "Sarah", 1).
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:   "case-2",
			output: model.Product{},
			err:    errors.Error("error"),
			mock: []interface{}{
				mock.ExpectExec("Update Products set Name=?,Type=? where Id=?").WithArgs("Sarah Vaughan", "Sarah", 0).
					WillReturnError(errors.Error("errpr")),
			},
		},
	}

	for _, tests := range tests {
		tc := tests
		store := New()

		output, err := store.UpdateByID(ctx, tc.output)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %v, got error: %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("expected %v, got :  %v", tc.output, output)
		}
	}
}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db

	tests := []struct {
		desc   string
		id     int
		output model.Product
		err    error
		mock   []interface{}
	}{
		{
			desc: "case-1",
			id:   1,
			err:  nil,
			mock: []interface{}{
				mock.ExpectExec("Delete from Products where id=?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc: "case-2",
			id:   0,
			err:  errors.Error("error"),
			mock: []interface{}{
				mock.ExpectExec("Update Products set Name=?,Type=? where Id=?").WithArgs(0).WillReturnError(errors.Error("error")),
			},
		},
	}
	for _, tests := range tests {
		tc := tests
		store := New()

		err := store.DeleteByID(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %v, got error: %v", tc.err, err)
		}
	}
}
