package product

import (
	"context"
	"testing"
	"fmt"
//	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	models "zopsmart/productgofr/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"

)


func TestGetProdById(t *testing.T) {
	app := gofr.New()

	db, mock, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	testCases := []struct {
		desc        string
		id          int
		expectedErr error
		Mock        []interface{}
		expectedRes *models.Product
	}{
		{
			desc: "Success case",
			id:   1,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
						AddRow(1, "shirt", "fashion")),
			},
			expectedErr: nil,
			expectedRes: &models.Product{
				Id:   1,
				Name: "shirt",
				Type: "fashion",
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
			expectedRes: &models.Product{},
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.GetProdByID(ctx,ts.id)
			if err != nil && ts.expectedErr != err {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}
		
	})
  }
}
