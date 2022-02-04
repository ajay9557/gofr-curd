package products

import (
	"context"
	"gofr-curd/models"
	"log"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	testProductsGetById(t, app)
	//TestProductDeleteById(t, app)
}

func testProductsGetById(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}
	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	query := "Select Id,Name,Type from Product where Id =?"
	tcs := []struct {
		desc           string
		Id             int
		err            error
		expectedOutput models.Product
		Mock           []interface{}
	}{
		{
			desc: "Success",
			Id:   3,
			err:  nil,
			expectedOutput: models.Product{
				Id:   3,
				Name: "Shirts",
				Type: "Twills",
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "Shirtspio", "US POLO"))},
		},
		{
			desc:           "Failure",
			Id:             0,
			err:            errors.EntityNotFound{Entity: "Product", ID: "0"},
			expectedOutput: models.Product{},
			Mock:           nil,
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.GetId(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}

// func TestProductDeleteById(t *testing.T) {
// 	app := gofr.New()
// 	tcs := []struct {
// 		desc string
// 		Id   int
// 		err  error
// 	}{
// 		{
// 			desc: "Success",
// 			Id:   1,
// 			err:  nil,
// 		},
// 		{
// 			desc: "Failure",
// 			Id:   4,
// 			err: errors.Error("Internal DB error"),
// 		},
// 	}

// 	for _, tc := range tcs {
// 		ctx := gofr.NewContext(nil, nil, app)
// 		ctx.Context = context.Background()
// 		store := New()
// 		t.Run(tc.desc, func(t *testing.T) {
// 			err := store.DeleteId(ctx, tc.Id)
// 			if !reflect.DeepEqual(err, tc.err) {
// 				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
// 			}
// 		})
// 	}
// }
