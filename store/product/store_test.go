package product

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/tejas/gofr-crud/models"
)

func TestGetProductById(t *testing.T) {

	app := gofr.New()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal("error in stub database connection")
	}

	defer db.Close()

	database, err :=  gorm.Open("mysql", db)

	if err != nil {
		fmt.Println("error in gorm database connection")
	}

	app.ORM = database

	testcases := []struct {
		desc        string
		input       int
		expectedErr error
		mockCall    []interface{}
		expected    models.Product
	}{
		{
			desc:        "Case 1: Success Case",
			input:       1,
			expectedErr: nil,
			mockCall: []interface{}{
				mock.ExpectQuery("select id, name, type from product where id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
				AddRow(1, "name-1", "test-1")),
			},
			expected: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
	},
		{
			desc:        "Case 2: Failure Case",
			input:       0,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "0"},
			mockCall: []interface{}{
				mock.ExpectQuery("select id, name, type from product").
					WithArgs(0).
					WillReturnError(errors.EntityNotFound{Entity: "product", ID: "0"}),
			},
			expected: models.Product{},
		},
}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.GetProductById(ctx, test.input)
			if err != nil && !reflect.DeepEqual(test.expectedErr, err) {
				fmt.Print("expected ", test.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(test.expected, res) {
				fmt.Print("expected ", test.expected, "obtained", res)
			}
		})
	}

}
