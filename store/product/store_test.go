package product

import (
	"context"
	"database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"reflect"
	"testing"
)

//func TestCoreLayer(t *testing.T) {
//	app := gofr.New()
//	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	//database, err := gorm.Open("mysql", db)
//	//if err != nil {
//	//	log.Print("error in opening gorm mysql")
//	//}
//	//app.ORM = database
//
//
//}

func TestProductStore_GetByID(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testcases := []struct {
		desc      string
		inp       int
		exp       *models.Product
		expErr    error
		mockCalls []interface{}
	}{
		{
			"success case",
			1,
			&models.Product{
				ID:   1,
				Name: "pavilion",
				Type: "laptop",
			},
			nil,
			[]interface{}{
				mock.ExpectQuery("select * from product where id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
						AddRow(1, "pavilion", "laptop")),
			},
		},
		{
			"failure case",
			1,
			nil,
			errors.EntityNotFound{
				Entity: "product",
				ID:     fmt.Sprint(1),
			},
			[]interface{}{
				mock.ExpectQuery("select * from product where id = ?").
					WithArgs(1).WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ctx.DB().DB = db
		store := New()
		out, err := store.GetByID(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestProductStore_GetAll(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testcases := []struct {
		desc      string
		exp       []*models.Product
		expErr    error
		mockCalls []interface{}
	}{
		{
			"success case",
			[]*models.Product{
				&models.Product{
					ID:   1,
					Name: "pavilion",
					Type: "laptop",
				}, &models.Product{
					ID:   2,
					Name: "testName",
					Type: "testType",
				}},
			nil,
			[]interface{}{
				mock.ExpectQuery("select id, name, type from product").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type"}).
						AddRow(1, "pavilion", "laptop").AddRow(2, "testName", "testType")),
			},
		},
		{
			"failure case",
			nil,
			errors.EntityNotFound{
				Entity: "product",
			},
			[]interface{}{
				mock.ExpectQuery("select id, name, type from product").
					WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ctx.DB().DB = db
		store := New()
		out, err := store.GetAll(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestProductStore_Delete(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testcases := []struct {
		desc      string
		inp       int
		expErr    error
		mockCalls []interface{}
	}{
		{
			"success case",
			1,
			nil,
			[]interface{}{
				mock.ExpectExec("DELETE FROM product where id = ?").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
		{
			"failure case",
			1,
			errors.EntityNotFound{
				Entity: "product",
				ID:     fmt.Sprint(1),
			},
			[]interface{}{
				mock.ExpectExec("DELETE FROM product where id = ?").
					WithArgs(1).WillReturnError(errors.EntityNotFound{
					Entity: "product",
					ID:     fmt.Sprint(1),
				}),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ctx.DB().DB = db
		store := New()
		err := store.Delete(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}

func TestProductStore_Create(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testcases := []struct {
		desc      string
		inp       *models.Product
		exp       *models.Product
		expErr    error
		mockCalls []interface{}
	}{
		{
			"success case",
			&models.Product{
				Name: "pavilion",
				Type: "laptop",
			},
			&models.Product{
				ID:   1,
				Name: "pavilion",
				Type: "laptop",
			},
			nil,
			[]interface{}{
				mock.ExpectExec("insert into product (name, type) values(?,?)").WithArgs("pavilion", "laptop").
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			"failure case",
			&models.Product{
				Name: "pavilion",
				Type: "laptop",
			},
			nil,
			errors.DB{Err: sql.ErrNoRows},
			[]interface{}{
				mock.ExpectExec("insert into product (name, type) values(?,?)").WithArgs("pavilion", "laptop").
					WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ctx.DB().DB = db
		store := New()
		out, err := store.Create(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestProductStore_Update(t *testing.T) {
	app := gofr.New()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	testcases := []struct {
		desc      string
		inp       *models.Product
		expErr    error
		mockCalls []interface{}
	}{
		{
			"update name",
			&models.Product{
				ID:   1,
				Name: "pavilion",
				Type: "",
			},
			nil,
			[]interface{}{
				mock.ExpectExec("update product set name = ? where id = ?").
					WithArgs("pavilion", 1).
					WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
		{
			"failure case",
			&models.Product{
				ID:   1,
				Name: "",
				Type: "laptop",
			},
			sql.ErrNoRows,
			[]interface{}{
				mock.ExpectExec("update product set type = ? where id = ?").
					WithArgs("laptop", 1).WillReturnError(sql.ErrNoRows),
			},
		},
	}

	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		ctx.DB().DB = db
		store := New()
		err := store.Update(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}
