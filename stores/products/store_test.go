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
	testInsertProduct(t, app)
	testUpdateProduct(t, app)
	testProductsGetById(t, app)
	testProductDeleteById(t, app)
	testAllProducts(t, app)
}

func testInsertProduct(t *testing.T, app *gofr.Gofr) {

	tcs := []struct {
		desc           string
		input          models.Product
		err            error
		expectedOutput models.Product
	}{
		{
			desc: "Success",
			err:  nil,
			input: models.Product{
				Id:   6,
				Name: "Brand",
				Type: "Mafti",
			},
			expectedOutput: models.Product{
				Id:   6,
				Name: "Brand",
				Type: "Mafti",
			},
		},
		{
			desc: "Failure",
			err:  errors.Error("Internal DB error"),
			input: models.Product{
				Id:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoih",
				Type: "dummy",
			},
			expectedOutput: models.Product{
				Id:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoih",
				Type: "dummy",
			},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.CreateProducts(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}

func testUpdateProduct(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc  string
		id    int
		err   error
		input models.Product
	}{
		{
			desc: "Success",
			id:   6,
			err:  nil,
			input: models.Product{
				Id:   6,
				Name: "Brands",
				Type: "Twills",
			},
		},
		{
			desc: "Failure case",
			id:   3,
			input: models.Product{
				Id:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoih",
				Type: "food",
			},
			err: errors.Error("Internal DB error"),
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := store.UpdateId(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
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
			Id:   6,
			err:  nil,
			expectedOutput: models.Product{
				Id:   6,
				Name: "Brands",
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
func testProductDeleteById(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc string
		Id   int
		err  error
	}{
		{
			desc: "Success",
			Id:   6,
			err:  nil,
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()
		t.Run(tc.desc, func(t *testing.T) {
			err := store.DeleteId(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testAllProducts(t *testing.T, app *gofr.Gofr) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	store := New()
	_, err := store.GetAll(ctx)
	if err != nil {
		t.Errorf("FAILED, Expected: %v, Got: %v", nil, err)
	}
}
