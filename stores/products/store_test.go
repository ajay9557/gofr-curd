package products

import (
	"context"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/arohanzst/testapp/models"

	perror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
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
						AddRow(1, "Biscuit", "Grocery")),
			},
			expOut: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
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
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"})),
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
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testRead(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	//data := models.Product{1, "Biscuit", }
	testCases := []struct {
		desc      string
		mockCalls []*sqlmock.ExpectedQuery
		expOut    []models.Product
		expErr    error
	}{
		{
			desc:   "Success",
			expErr: nil,
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT Id, Name, Type FROM Product").
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "Biscuit", "Grocery")),
			},
			expOut: []models.Product{
				{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				},
			},
		},
		{
			desc:   "Failure",
			expErr: errors.New("Error in Given Query"),
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT Id, Name, Type FROM Product").
					WithArgs().
					WillReturnError(errors.New("Error in Given Query")),
			},
			expOut: nil,
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
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testCreate(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	//data := models.Product{1, "Biscuit", }
	testCases := []struct {
		desc   string
		input  models.Product
		expErr error
		expOut *models.Product
	}{
		{
			desc:   "Test Case 1",
			input:  models.Product{Name: "Trimmer", Type: "Electric"},
			expErr: nil,
			expOut: &models.Product{Id: 17, Name: "Trimmer", Type: "Electric"},
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
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testUpdate(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	testCases := []struct {
		desc   string
		input  models.Product
		id     int
		expErr error
		expOut *models.Product
	}{
		{
			desc:   "Test Case 1",
			input:  models.Product{Name: "Study-Lamp", Type: "Electric"},
			id:     17,
			expErr: nil,
			expOut: &models.Product{Id: 17, Name: "Study-Lamp", Type: "Electric"},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.Update(ctx, &tc.input, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testDelete(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	testCases := []struct {
		desc   string
		id     int
		expErr error
		expOut *models.Product
	}{
		{
			desc:   "Test Case 1",
			id:     17,
			expErr: nil,
			expOut: nil,
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
