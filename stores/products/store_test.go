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
	testProductsGetByID(t, app)
	testProductDeleteByID(t, app)
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
				ID:   6,
				Name: "Brand",
				Type: "Mafti",
			},
			expectedOutput: models.Product{
				ID:   6,
				Name: "Brand",
				Type: "Mafti",
			},
		},
		{
			desc: "Failure",
			err:  errors.Error("Internal DB error"),
			input: models.Product{
				ID:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoia",
				Type: "dummy",
			},
			expectedOutput: models.Product{
				ID:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoia",
				Type: "dummy",
			},
		},
	}

	for _, test := range tcs {
		tc := test
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
		ID    int
		err   error
		input models.Product
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
			input: models.Product{
				ID:   6,
				Name: "Brands",
				Type: "Twills",
			},
		},
		{
			desc: "Failure case",
			ID:   3,
			input: models.Product{
				ID:   3,
				Name: "very-long-namebviuauefieufohoiahhwoieflruogeroigruigwuoehfihoweinveoihvery-long-namebviuauefieufohoia",
				Type: "food",
			},
			err: errors.Error("Internal DB error"),
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			err := store.UpdateID(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func testProductsGetByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}

	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	defer db.Close()

	database, err := gorm.Open("mysql", db)
	if err != nil {
		log.Println("Error opening gorm conn", db)
	}

	app.ORM = database
	query := "Select ID,Name,Type from Product where ID =?"
	tcs := []struct {
		desc           string
		ID             int
		err            error
		expectedOutput models.Product
		Mock           []interface{}
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
			expectedOutput: models.Product{
				ID:   6,
				Name: "Brands",
				Type: "Twills",
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows(
				[]string{"ID", "Name", "Type"}).AddRow(1, "Shirtspio", "US POLO"))},
		},
		{
			desc:           "Failure",
			ID:             0,
			err:            errors.EntityNotFound{Entity: "Product", ID: "0"},
			expectedOutput: models.Product{},
			Mock:           nil,
		},
	}
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	store := New()

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			res, err := store.GetID(ctx, tc.ID)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}
func testProductDeleteByID(t *testing.T, app *gofr.Gofr) {
	tcs := []struct {
		desc string
		ID   int
		err  error
	}{
		{
			desc: "Success",
			ID:   6,
			err:  nil,
		},
	}

	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		store := New()

		t.Run(tc.desc, func(t *testing.T) {
			err := store.DeleteID(ctx, tc.ID)
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
