package products

import (
	"context"
	"log"
	"reflect"
	"testing"

	perror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arohanzst/testapp/models"
	"github.com/jinzhu/gorm"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	database, err := gorm.Open("mysql", db)

	if err != nil {
		log.Println("Error in opening gorm conn", db)
	}

	app.ORM = database

	testReadByID(t, app, mock)
	testRead(t, app, mock)
	testCreate(t, app, mock)
	testUpdate(t, app, mock)
	testDelete(t, app, mock)
}

func testReadByID(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	testCases := []struct {
		desc      string
		input     int
		mockCalls []*sqlmock.ExpectedQuery
		expOut    *models.Product
		expErr    error
	}{
		{
			desc:   "Success",
			input:  1,
			expErr: nil,
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT Id, Name, Type FROM Product where Id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "Biscuit", "Daily-Use")),
			},
			expOut: &models.Product{
				ID:   1,
				Name: "Biscuit",
				Type: "Daily-Use",
			},
		},
		{
			desc:  "Failure: Product entity not present in Database",
			input: 10,
			expErr: perror.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT Id, Name, Type FROM Product where Id = ?").
					WithArgs(10).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "Biscuit", "Daily-Use")),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.ReadByID(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil {
			if !reflect.DeepEqual(out, tc.expOut) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
			}
		}
	}
}

func testRead(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	testCases := []struct {
		desc      string
		expOut    []models.Product
		expErr    error
		mockCalls []*sqlmock.ExpectedQuery
	}{
		{
			desc:   "Success",
			expErr: nil,
			expOut: []models.Product{
				{
					ID:   1,
					Name: "Biscuit",
					Type: "Daily-Use",
				},
			},
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT Id, Name, Type FROM Product").
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"})),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.Read(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil {
			if !reflect.DeepEqual(out, tc.expOut) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
			}
		}
	}
}

func testCreate(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	query := "INSERT INTO Product(Name, Type) values(?, ?)"
	testCases := []struct {
		desc      string
		input     models.Product
		expErr    error
		expOut    *models.Product
		mockCalls []*sqlmock.ExpectedExec
	}{
		{
			desc:      "Case 1: Success",
			input:     models.Product{Name: "Trimmer", Type: "Electric"},
			expErr:    nil,
			expOut:    &models.Product{ID: 27, Name: "Trimmer", Type: "Electric"},
			mockCalls: []*sqlmock.ExpectedExec{mock.ExpectExec(query).WithArgs("Trimmer", "Electric").WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.Create(ctx, &tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil {
			if !reflect.DeepEqual(out, tc.expOut) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
			}
		}
	}
}

func testUpdate(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	query, _ := MakeUpdateQuery(&models.Product{Type: "Daily-Use"}, 23)

	testCases := []struct {
		desc      string
		input     models.Product
		id        int
		expErr    error
		expOut    *models.Product
		mockCalls []*sqlmock.ExpectedExec
	}{
		{
			desc:      "Case 1: Success",
			input:     models.Product{Type: "Daily-Use"},
			id:        23,
			expErr:    nil,
			expOut:    &models.Product{ID: 23, Name: "Shampoo", Type: "Daily-Use"},
			mockCalls: []*sqlmock.ExpectedExec{mock.ExpectExec(query).WithArgs(23).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tc := range testCases {
		store := New()

		out, err := store.Update(ctx, &tc.input, tc.id)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil {
			if !reflect.DeepEqual(out, tc.expOut) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
			}
		}
	}
}

func testDelete(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	query := "DELETE FROM Product where Id=?"
	testCases := []struct {
		desc      string
		id        int
		expErr    error
		mockCalls []*sqlmock.ExpectedExec
		expOut    *models.Product
	}{
		{
			desc:      "Case 1: Success",
			id:        23,
			expErr:    nil,
			expOut:    nil,
			mockCalls: []*sqlmock.ExpectedExec{mock.ExpectExec(query).WithArgs(23).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		err := store.Delete(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}
