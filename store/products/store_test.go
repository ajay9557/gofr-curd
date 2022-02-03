package products

import (
	"context"
	"log"
	"reflect"
	"testing"
	"zopsmart/gofr-curd/model"

	"github.com/jinzhu/gorm"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	defer db.Close()
	rows := sqlmock.NewRows([]string{"Id", "Name", "Type"}).
		AddRow(1, "Reebok", "Bats")
	tests := []struct {
		desc   string
		id     int
		err    error
		output model.Product
		mock   []interface{}
	}{
		{
			desc: "Get existent id",
			id:   1,
			err:  nil,
			output: model.Product{
				Id:   1,
				Name: "Reebok",
				Type: "Bats",
			},
			mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM User WHERE id = ?").WithArgs(1).WillReturnRows(rows),
			},
		},
		{
			desc: "Get existent id",
			id:   1223,
			err:  errors.EntityNotFound{Entity: "product", ID: "1223"},
			mock: nil,
		},
	}
	for _, tc := range tests {
		app := gofr.New()

		app.ORM = database
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()
		res, err := store.GetProductById(ctx, tc.id)
		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("expected %v got %v", tc.output, res)
		}
		if tc.err != err {
			t.Errorf("expected %s got %s", tc.err, err)
		}
	}
}
