package product

import (
	"context"
	"gofr-curd/models"
	"log"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}
	app.ORM = database
	testGetById(t, app, mock)
}

func testGetById(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	testCases := []struct {
		desc      string
		input     int
		mockCalls []*sqlmock.ExpectedQuery
		expOut    *models.Product
		expErr    error
	}{
		{
			desc:   "success case",
			input:  1,
			expErr: nil,
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("select name, type from products where id=?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
						AddRow(1, "test", "example")),
			},
			expOut: &models.Product{
				Id:   1,
				Name: "test",
				Type: "example",
			},
		},
		{
			desc:  "entity not in database",
			input: 1022,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1022",
			},
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("select name, type from products where id=?").
					WithArgs(1022).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"})),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		out, err := store.GetById(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}
