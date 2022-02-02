package products

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	db, mock, _ := sqlmock.New()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database

	testGetProductByID(t, app, mock)
}

func testGetProductByID(t *testing.T, app *gofr.Gofr, mock sqlmock.Sqlmock) {

	rows := mock.NewRows([]string{"id", "name", "category"}).AddRow(1, "mouse", "electronics")

	tests := []struct {
		desc     string
		id       int
		err      error
		mockCall *sqlmock.ExpectedQuery
	}{
		{
			desc:     "Get existent id",
			id:       1,
			err:      nil,
			mockCall: mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows),
		},
		{
			desc:     "Get non existent id",
			id:       100,
			err:      errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall: mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(100).WillReturnRows(mock.NewRows([]string{"id", "name", "category"})),
		},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			_, err := store.GetById(ctx, tc.id)
			assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}
