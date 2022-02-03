package product

import (
	"context"
	"database/sql"
	"log"
	"productGofr/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	// seeder := datastore.NewSeeder(&app.DataStore,"../db")
	// seeder.ResetCounter = true
	db, mock, _ := sqlmock.New()
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Can't Open the DataBase")
	}
	app.ORM = database

	rows := sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "daikinn", "AC")

	tests := []struct {
		desc        string
		id          int
		expected    *models.Product
		expectedErr error
		mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, expectedErr: nil, expected: &models.Product{Id: 1, Name: "daikinn", Type: "AC"}, mockQuery: mock.ExpectQuery("select * from Product where id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", id: 100, expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil, mockQuery: mock.ExpectQuery("select * from Product where id = ?").WithArgs(100).WillReturnError(sql.ErrNoRows)},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			res, err := istore.UserById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, res) {
				t.Error("expected: ", test.expected, "obtained: ", res)
			}
		})
	}

}
