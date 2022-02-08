package product

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gofr-curd/models"
	"gofr-curd/store"
	"log"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func setMock() (*gofr.Context, store.Store, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
	}

	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	ctx.DB().DB = db
	s := New()

	return ctx, s, mock, db
}

func TestCreate(t *testing.T) {
	ctx, s, mock, db := setMock()
	defer db.Close()

	tesCases := []struct {
		desc     string
		input    models.Product
		expErr   error
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc: "success case",
			input: models.Product{
				ID:   3,
				Name: "this",
				Type: "that",
			},
			expErr: nil,
			mockCall: mock.ExpectExec("insert into products(id,name,type) values(?,?,?)").
				WithArgs(3, "this", "that").WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc: "error case",
			input: models.Product{
				ID:   3,
				Name: "hello",
				Type: "moto",
			},
			expErr: errors.EntityAlreadyExists{},
			mockCall: mock.ExpectExec("insert into products(id,name,type) values(?,?,?)").
				WithArgs(3, "hello", "moto").WillReturnError(errors.EntityAlreadyExists{}),
		},
	}

	for _, test := range tesCases {
		tc := test
		err := s.Create(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}

}

func TestGet(t *testing.T) {
	ctx, s, mock, db := setMock()
	defer db.Close()

	testCases := []struct {
		desc     string
		expErr   error
		expOut   []*models.Product
		mockCall *sqlmock.ExpectedQuery
	}{
		{
			desc:   "success case",
			expErr: nil,
			expOut: []*models.Product{
				{
					ID:   1,
					Name: "test",
					Type: "example",
				},
				{
					ID:   2,
					Name: "this",
					Type: "that",
				},
			},
			mockCall: mock.ExpectQuery("select id,name,type from products").
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
					AddRow(1, "test", "example").
					AddRow(2, "this", "that")),
		},
		{
			desc:   "error entity not found",
			expErr: errors.EntityNotFound{Entity: "products", ID: "all"},
			expOut: nil,
			mockCall: mock.ExpectQuery("select id,name,type from products").
				WillReturnError(sql.ErrNoRows),
		},
		{
			desc:   "error scanning",
			expOut: nil,
			expErr: errors.EntityNotFound{Entity: "product"},
			mockCall: mock.ExpectQuery("select id,name,type from products").
				WillReturnRows(sqlmock.NewRows([]string{"name", "type"}).
					AddRow("test", "example")),
		},
	}
	for _, test := range testCases {
		tc := test
		out, err := s.Get(ctx)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func TestGetByID(t *testing.T) {
	ctx, s, mock, db := setMock()
	defer db.Close()

	testCases := []struct {
		desc     string
		input    int
		expErr   error
		expOut   *models.Product
		mockCall *sqlmock.ExpectedQuery
	}{
		{
			desc:   "success case",
			input:  1,
			expErr: nil,
			expOut: &models.Product{
				ID:   1,
				Name: "test",
				Type: "example",
			},
			mockCall: mock.ExpectQuery("select name, type from products where id=?").
				WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name", "type"}).
				AddRow("test", "example")),
		},
		{
			desc:  "entity not in database",
			input: 1022,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1022",
			},
			mockCall: mock.ExpectQuery("select name, type from products where id=?").
				WithArgs(1022).WillReturnError(sql.ErrNoRows),
		},
	}

	for _, test := range testCases {
		tc := test

		out, err := s.GetByID(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

type testCase struct {
	desc     string
	input    models.Product
	expErr   error
	mockCall *sqlmock.ExpectedExec
}

func TestUpdate(t *testing.T) {
	ctx, s, mock, db := setMock()
	defer db.Close()

	tesCases := []testCase{
		{
			desc: "success case",
			input: models.Product{
				ID:   3,
				Name: "hello",
				Type: "world",
			},
			expErr: nil,
			mockCall: mock.ExpectExec("update products set name=?, type=? where id=?").
				WithArgs("hello", "world", 3).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc: "error updating",
			input: models.Product{
				ID:   3,
				Name: "hello",
				Type: "world",
			},
			expErr: errors.Error("error updating record"),
			mockCall: mock.ExpectExec("update products set name=?, type=? where id=?").
				WithArgs("hello", "world", 3).WillReturnError(sql.ErrConnDone),
		},
	}

	for _, test := range tesCases {
		tc := test
		err := s.Update(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

func TestDelete(t *testing.T) {
	ctx, s, mock, db := setMock()
	defer db.Close()

	tesCases := []struct {
		desc     string
		input    int
		expErr   error
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:   "success case",
			input:  3,
			expErr: nil,
			mockCall: mock.ExpectExec("delete from products where id=?").
				WithArgs(3).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:   "error deleting record",
			input:  1022,
			expErr: errors.Error("error deleting record"),
			mockCall: mock.ExpectExec("delete from products where id=?").
				WithArgs(1022).WillReturnError(sql.ErrConnDone),
		},
	}

	for _, tc := range tesCases {
		err := s.Delete(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}
