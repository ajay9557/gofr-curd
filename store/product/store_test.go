package product

import (
	"context"
	"fmt"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestGetById(t *testing.T) {

	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		fmt.Println("Error opening gorm conn", db)
	}
	app.ORM = database

	testCases := []struct {
		desc        string
		id          int
		expectedErr error
		Mock        []interface{}
		expectedRes models.Product
	}{
		{
			desc: "Success case",
			id:   1,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "jeans", "clothes")),
			},
			expectedErr: nil,
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
		},
		{
			desc: "Failure case",
			id:   0,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1).
					WillReturnError(errors.EntityNotFound{Entity: "product", ID: "0"}),
			},
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "0"},
			expectedRes: models.Product{},
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.GetById(ts.id, ctx)
			if err != nil && !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}
		})
	}

}
