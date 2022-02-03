package products

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/arohanzst/testapp/models"

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
		log.Println("Error in opening gorm conn", db)
	}
	app.ORM = database
	testReadByID(t, app, mock)
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
				mock.ExpectQuery("SELECT * FROM Product where Id = ?").
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
			expErr: errors.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			mockCalls: []*sqlmock.ExpectedQuery{
				mock.ExpectQuery("SELECT * FROM Product where Id = ?").
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
