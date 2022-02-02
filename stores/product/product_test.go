package product

import (
	"context"
	"fmt"
	"log"
	"product/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func Test_GetProductById(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	mock.ExpectQuery("......")
	testCases := []struct {
		desc           string
		id             int
		expectedError  error
		mock           []interface{}
		expectedResult models.Product
	}{
		{
			desc:          "Test Case 1",
			id:            2,
			expectedError: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from product where id=?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "jeans", "clothes")),
			},
		},
		{
			desc:          "Test Case 2",
			id:            0,
			expectedError: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from product where id=?").WithArgs(0).WillReturnError(errors.EntityNotFound{Entity: "product", ID: "0"}),
			},
		},
	}

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.GetProductById(ctx, tcs.id)
			if err != nil && !reflect.DeepEqual(tcs.expectedError, err) {
				fmt.Print("expected ", tcs.expectedError, "obtained", err)
			}
			if !reflect.DeepEqual(tcs.expectedError, res) {
				fmt.Print("expected ", tcs.expectedError, "obtained", res)
			}
		})
	}
}
