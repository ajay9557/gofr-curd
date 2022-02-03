package PRODUCT

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	gofrerr "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

func testGetById(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New()
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database

	mock.ExpectQuery("......")

	testcase := []struct {
		desc   string
		id     int
		experr error
		mock   []interface{}
		expout *model.Product
	}{
		{
			desc:   "test case sucess",
			id:     1,
			experr: nil,
			expout: &model.Product{
				Id:   1,
				Name: "tatakai",
				Type: "koros",
			},
			mock: []interface{}{
				mock.ExpectQuery("select  Name, Type FROM user where Id=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "tatakai", "koros")),
			},
		},
		{
			desc: "Test case 2",
			id:   0,
			experr: gofrerr.EntityNotFound{
				Entity: "user",
				ID:     "0",
			},
			mock: []interface{}{
				mock.ExpectQuery("select Name,Type from user where id=?").WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"})),
			},
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			sto := New()
			res, err := sto.GetById(tcs.id, ctx)
			if err != nil && !errors.Is(tcs.experr, err) {
				fmt.Print("expected ", tcs.experr, "obtained", err)
			}
			if !reflect.DeepEqual(tcs.experr, res) {
				fmt.Print("expected ", tcs.experr, "obtained", res)
			}
		})
	}

}
