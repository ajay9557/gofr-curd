package product

import (
	"context"
	"database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"testing"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Print("error in opening gorm mysql")
	}
	app.ORM = database
	testGetByID(t, app, mock)

}

func testGetByID(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {
	testcases := []struct {
		desc      string
		inp       int
		exp       *models.Product
		expErr    error
		mockCalls []interface{}
	}{
		{
			"success case",
			1,
			&models.Product{
				Id:   1,
				Name: "legion",
				Type: "laptop",
			},
			nil,
			[]interface{}{
				mock.ExpectQuery("select * from product where id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
						AddRow(1, "legion", "laptop")),
			},
		},
		{
			"failure case",
			100,
			nil,
			errors.EntityNotFound{
				Entity: "product",
				ID:     fmt.Sprint(100),
			},
			[]interface{}{
				mock.ExpectQuery("select * from product where id = ?").
					WithArgs(100).WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		out, err := store.GetByID(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
