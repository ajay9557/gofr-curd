package product

import (
	"context"
	"database/sql"
	"fmt"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStoreLayer(t *testing.T) {
	app := gofr.New()

	testGetProductByID(t, app)
	testGetAllProducts(t, app)
	testCreateProduct(t, app)
	testDeleteProduct(t, app)
	testUpdateProductByID(t, app)
}
func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "type"}).AddRow(
		1, "daikinn", "AC",
	)

	defer db.Close()

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc        string
		id          int
		expects     *models.Product
		expectedErr error
		mockQuery   *sqlmock.ExpectedQuery
	}{
		{
			desc:        "Case1",
			id:          1,
			expectedErr: nil,
			expects:     &models.Product{Id: 1, Name: "daikinn", Type: "AC"},
			mockQuery:   mock.ExpectQuery("select * from Product where id = ?").WithArgs(1).WillReturnRows(rows)},
		{
			desc:        "Case2",
			id:          100,
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			expects:     nil,
			mockQuery:   mock.ExpectQuery("select * from Product where id = ?").WithArgs(100).WillReturnError(sql.ErrNoRows)},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			istore := New()
			ctx.DB().DB = db
			res, err := istore.GetProductByID(ctx, test.id)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expects, res) {
				t.Error("expected: ", test.expects, "obtained: ", res)
			}
		})
	}
}

func testGetAllProducts(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "type"}).
		AddRow(1, "daikinn", "AC").AddRow(2, "milton", "Water Bottle")

	defer db.Close()

	tests := []struct {
		desc string
		// id          int
		expects     []*models.Product
		expectedErr error
		mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1",
			expectedErr: nil,
			expects: []*models.Product{{Id: 1, Name: "daikinn", Type: "AC"},
				{Id: 2, Name: "milton", Type: "Water Bottle"},
				// &models.Product{Id: 3, Name: "Kenstar", Type: "Microwave"},
				// &models.Product{Id: 4, Name: "Ultra", Type: "RedGrinder"},
				// &models.Product{Id: 5, Name: "Crompton", Type: "Fan"},
				// &models.Product{Id: 6, Name: "Prestige", Type: "RiceCooker"},
				// &models.Product{Id: 13, Name: "Nivvea", Type: "Moisturizzerr"},
				// &models.Product{Id: 16, Name: "Kenstarr1", Type: "Microwavee1"},
				// &models.Product{Id: 18, Name: "Kenstarr7", Type: "Microwavee7"},
				// &models.Product{Id: 21, Name: "daikin", Type: "AC"},
			},
			mockQuery: mock.ExpectQuery("select * from Product ").WillReturnRows(rows),
		},
		{desc: "Case2",
			expectedErr: errors.Error("Couldn't execute query"),
			expects:     []*models.Product{},
			mockQuery:   mock.ExpectQuery("select * from Product ").WillReturnError(errors.Error("Couldn't execute query")),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			ctx.DB().DB = db
			res, err := istore.GetAllProducts(ctx)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expects, res) {
				t.Error("expected: ", test.expects, "obtained: ", res)
			}
		})
	}
}
func testCreateProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	tests := []struct {
		desc string
		// id          int
		input models.Product
		// expected    *models.Product
		expects     int
		expectedErr error
		mockQuery   *sqlmock.ExpectedExec
	}{
		{desc: "Case1", /*id: 1,*/
			input:       models.Product{Name: "daikin", Type: "AC"},
			expectedErr: nil,
			expects:     1,
			mockQuery: mock.ExpectExec("insert into Product(name,type) values (?,?)").
				WithArgs("daikin", "AC").WillReturnResult(sqlmock.NewResult(1, 1))},

		{desc: "Case2", /*id: 100,*/
			input:       models.Product{Name: "daikin", Type: "AC"},
			expectedErr: errors.Error("Couldn't execute query"),
			expects:     0,
			mockQuery: mock.ExpectExec("insert into Product(name,type) values (?,?)").
				WithArgs("daikin", "AC").WillReturnError(errors.Error("Couldn't execute query"))},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			ctx.DB().DB = db
			res, err := istore.CreateProduct(ctx, test.input)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expects, res) {
				t.Error("expected: ", test.expects, "obtained: ", res)
			}
		})
	}
}
func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tests := []struct {
		desc string
		id   int
		// expected    *models.Product
		expectedErr error
		mockQuery   *sqlmock.ExpectedExec
	}{
		{desc: "Case1",
			id: 1, expectedErr: nil,
			mockQuery: mock.ExpectExec("delete from Product where id = ?").
				WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{desc: "Case2", id: 100, expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockQuery: mock.ExpectExec("delete from Product where id = ?").
				WithArgs(100).WillReturnResult(sqlmock.NewResult(0, 0)),
		},
		{desc: "Case3",
			id: 2, expectedErr: errors.Error("Couldn't execute query"),
			mockQuery: mock.ExpectExec("delete from Product where id = ?").
				WithArgs(1).WillReturnError(errors.Error("Couldn't execute query")),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			ctx.DB().DB = db
			err := istore.DeleteByID(ctx, test.id)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}
func testUpdateProductByID(t *testing.T, app *gofr.Gofr) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tests := []struct {
		desc    string
		id      int
		input   models.Product
		expects int
		// expected    int
		expectedErr error
		mockQuery   *sqlmock.ExpectedExec
	}{
		{desc: "Case1", id: 21, input: models.Product{Name: "daikinn", Type: "ACC"},
			expectedErr: nil, expects: 21,
			mockQuery: mock.ExpectExec("update Product set name = ?, type = ? where id = ?").
				WithArgs("daikinn", "ACC", 21).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{desc: "Case2", id: 21, input: models.Product{},
			expectedErr: errors.Error("Nothing to Update"), expects: 0,
			mockQuery: nil},
		{desc: "Case3", id: 21, input: models.Product{Name: "daikin", Type: "AC"},
			expectedErr: errors.Error("Couldn't execute query"), expects: 0,
			mockQuery: mock.ExpectExec("update Product set name = ?, type = ? where id = ?").
				WithArgs("daikin", "AC", 21).WillReturnError(errors.Error("Couldn't execute query")),
		},
		{desc: "Case4", id: 21, input: models.Product{Name: "daikin", Type: "AC"},
			expectedErr: errors.Error("SAME DATA GIVEN TO PREVIOUS DATA"), expects: 21,
			mockQuery: mock.ExpectExec("update Product set name = ?, type = ? where id = ?").
				WithArgs("daikin", "AC", 21).WillReturnResult(sqlmock.NewResult(0, 0)),
		},

		// {desc: "Case2", /*id: 100,*/ input:,expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			ctx.DB().DB = db
			res, err := istore.UpdateByID(ctx, test.id, test.input)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expects, res) {
				t.Error("expected: ", test.expects, "obtained: ", res)
			}
		})
	}
}
