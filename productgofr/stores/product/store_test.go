package product

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	models "zopsmart/productgofr/models"

	//	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
//	"golang.org/x/tools/go/expect"

	//	"github.com/modern-go/reflect2"

	//	"gorm.io/driver/mysql"
	//	"gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

func TestGetProdById(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	database, err := gorm.Open("mysql",db)
	if err != nil {
		t.Error(err)
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
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(ts.desc, func(t *testing.T) {
			res, err := store.GetProdByID(ctx, ts.id)
			if err != nil && ts.expectedErr != err {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}

		})
	}
}


func TestCreateProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	database, err := gorm.Open("mysql",db)
	if err != nil {
		t.Error(err)
	}

	app.ORM = database


	testCases := []struct {
		desc        string
		input *models.Product
		expectedErr error
		Mock        []interface{}
		expectedRes *models.Product
	}{
		{
			desc: "Success case",
			input:   &models.Product{
				Id: 1,
				Name:  "shirt",
				Type:  "fashion",
			},
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(models.Product{
					Id :1,
					Name: "shirt",
					Type: "fashion",
				}).
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
			desc: "Failure case 1",
			input:  &models.Product{} ,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(models.Product{}).
					WillReturnError(errors.DB{Err:err}),
			},
			expectedErr: errors.DB{Err:err},
			expectedRes: &models.Product{},
		},
		{
			desc: "Failure case 2",
			input:  &models.Product{} ,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(models.Product{
				Id:  0,
				Name: "shirt",
				Type: "fashion",		
				}).
					WillReturnError(errors.DB{Err:err}),
			},
			expectedErr: errors.DB{Err:err},
			expectedRes: &models.Product{
				Id:  0,
				Name: "shirt",
				Type: "fashion",
			},
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res, err := store.CreateProduct(ctx, ts.input)
			if err != nil && ts.expectedErr != err {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}

		})
	}
	

}

func TestDeleteProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	database, err := gorm.Open("mysql",db)
	if err != nil {
		t.Error(err)
	}

	app.ORM = database
	testCases := []struct {
		desc        string
		input int
		expectedErr error
		Mock        []interface{}
	} {
		{
			desc: "Success Case",
			input: 1,
			expectedErr: nil,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(1).
					WillReturnError(nil)},
		},
		{
			desc: "Failure Case",
			input: 443,
			expectedErr: errors.DB{Err: err},
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Type from product").WithArgs(443).
					WillReturnError(errors.DB{Err: err})},
		},
	}

	for _,ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		err := store.DeleteProduct(ctx,ts.input)

		if !reflect.DeepEqual(ts.expectedErr,err) {
			fmt.Print("expected ", ts.expectedErr, "obtained", err)
		}

		})
	}
}

func TestUpdateProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	database, err := gorm.Open("mysql",db)
	if err != nil {
		t.Error(err)
	}

	app.ORM = database

	testcase := []struct {
		desc string
		input models.Product
		expectedRes *models.Product
		expectedErr error
	}{
		{
			desc: "Success case",
			input: models.Product{
				Id:1,
				Name: "fruits",
				Type: "Daily needs",
			},
			expectedRes: &models.Product{
				Id:1,
				Name: "fruits",
				Type: "Daily needs",
			},
			expectedErr: nil,

		},
		{
			desc: "Failure Case",
			input: models.Product{},
			expectedRes: &models.Product{},
			expectedErr: errors.DB{Err:err},

		},
	}

	for _,ts := range testcase {
		t.Run(ts.desc, func(t *testing.T) {
			query := "UPDATE product SET"
			fields,values := formQuery(ts.input)
			query += fields + " WHERE id = ?"
			mock.ExpectQuery(query).WithArgs(values[0], values[1],values[2]).WillReturnError(nil)
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			store := New()
			res,err := store.UpdateProduct(ctx,ts.input)
			if !reflect.DeepEqual(err,ts.expectedErr) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}

			if !reflect.DeepEqual(res,ts.expectedRes) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}

		})
	}

}

func TestGetAllProduct(t *testing.T) {
	app := gofr.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	database, err := gorm.Open("mysql",db)
	if err != nil {
		t.Error(err)
	}

	app.ORM = database


	testCases := []struct {
		expectedErr error
		expectedOut []*models.Product
		mockQuery *sqlmock.ExpectedQuery
	} {
		{
			expectedErr: nil,
			mockQuery: mock.ExpectQuery("SELECT Id,Name,Type from product").
				WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Type"}).
					AddRow(1, "shirt","fashion").
					AddRow(2, "mobile", "electronics")),
			expectedOut: []*models.Product{
				{
					Id:    1,
					Name:  "shirt",
					Type: "fashion",
				},
				{
					Id:    2,
					Name:  "mobile",
					Type: "electronics",
				},
			},
		},
	}
    for _,tc := range testCases {
		t.Run("testing",func(t *testing.T) {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		res, err := store.GetAllProduct(ctx)
		if !reflect.DeepEqual(res, tc.expectedOut) {
			t.Errorf("Expected: \t%v\nGot: \t%v\n",tc.expectedOut, res)
		}
		if !reflect.DeepEqual(err,tc.expectedErr) {
			t.Errorf("Expected: \t%v\nGot: \t%v\n",tc.expectedErr, err)
		}
	})
			
	}

}
